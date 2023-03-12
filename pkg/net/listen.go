package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Listen 底层是一个Listen状态的socket
	ln, err := net.Listen("tcp", ":8888")

	if err != nil {
		panic(err)
	}

	conn, err := ln.Accept()

	if err != nil {
		panic(err)
	}

	var body [100]byte
	for true {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到消息：%s\n", string(body[:]))
		conn.Write(body[:])
		if err != nil {
			break
		}
	}
}
