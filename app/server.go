package main

import (
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
	request := newRequest(conn)
	if request.path == "/" {
		EmptyOkResponse().send(conn)
	} else if strings.HasPrefix(request.path, "/echo") {
		data := strings.Split(request.path, "/echo/")[1]
		OkResponse(data).send(conn)
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
