package main

// TCP connection

import (
	"fmt"
	"io"
	"net"
)

var conx_match = make(chan io.ReadWriteCloser)

func main() {
	listener, err := net.Listen("tcp", "localhost:8181")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conx, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go pair(conx)
	}
}

func pair(cx1 io.ReadWriteCloser) {
	fmt.Println("Peer connection has entered")
	select {
	case conx_match <- cx1:
		// Handled by goroutine
	case cx2 := <-conx_match:
		sendchat(cx2, cx1)
	}
}

func sendchat(a, b io.ReadWriteCloser) {
	go io.Copy(a, b)
	io.Copy(b, a)
}
