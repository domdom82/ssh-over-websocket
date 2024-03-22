package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
	"sync"
)

func main() {

	url := os.Args[1]

	wsConn, err := websocket.Dial(url, "", url)

	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		_, err := io.Copy(wsConn, os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
		}
		wg.Done()
	}()

	go func() {
		_, err := io.Copy(os.Stdout, wsConn)
		if err != nil {
			fmt.Println(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
	_ = wsConn.Close()
}
