package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("클라이언트 연결 종료:", err)
			return
		}
		fmt.Printf("받은 메시지: %s", message)
		conn.Write([]byte(message))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("서버 시작: 포트 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("연결 수락 실패:", err)
			continue
		}
		go handleConnection(conn)
	}
}
