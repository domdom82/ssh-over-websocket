package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {

	mux := http.NewServeMux()

	ws := websocket.Server{
		Handshake: nil,
		Handler:   forwardToSSH,
	}

	mux.Handle("/ws", ws)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func forwardToSSH(wsConn *websocket.Conn) {
	defer func() {
		_ = wsConn.Close()
	}()

	sshConn, err := net.Dial("tcp", "localhost:22")

	if err != nil {
		log.Println("Could not connect to ssh backend:", err.Error())
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		_, err := io.Copy(sshConn, wsConn)
		if err != nil {
			log.Println(err.Error())
		}
		wg.Done()
	}()

	go func() {
		_, err := io.Copy(wsConn, sshConn)
		if err != nil {
			log.Println(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
