package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

func main() {

	url := os.Args[1]

	wsConn, err := websocket.Dial(url, "", url)

	if err != nil {
		panic(err)
	}

	sshClient, err := wrapSSH(wsConn, url)

	if err != nil {
		panic(err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func wrapSSH(conn *websocket.Conn, addr string) (*ssh.Client, error) {
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: "testuser",
		Auth: []ssh.AuthMethod{
			ssh.Password("yourpassword"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, addr, config)
	if err != nil {
		return nil, err
	}
	return ssh.NewClient(c, chans, reqs), nil
}
