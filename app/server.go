package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	path := getRequestedPath(conn)
	if path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}

func getRequestedPath(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	startLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading start line: ", err.Error())
		os.Exit(1)
	}
	return strings.Split(startLine, " ")[1]
}
