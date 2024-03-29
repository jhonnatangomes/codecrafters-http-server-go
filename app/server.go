package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"strings"
)

var dir = flag.String("directory", ".", "directory to serve")

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	request := newRequest(conn)
	switch {
	case request.path == "/":
		EmptyOkResponse().send(conn)
	case strings.HasPrefix(request.path, "/echo"):
		data := strings.Split(request.path, "/echo/")[1]
		OkResponse([]byte(data)).send(conn)
	case request.path == "/user-agent":
		OkResponse([]byte(request.headers["User-Agent"])).send(conn)
	case strings.HasPrefix(request.path, "/files"):
		filename := strings.Split(request.path, "/files/")[1]
		if request.method == "GET" {
			if file, err := os.ReadFile(path.Join(*dir, filename)); err == nil {
				FileOkResponse(file).send(conn)
			} else if os.IsNotExist(err) {
				NotFoundResponse().send(conn)
			} else {
				fmt.Println("Error reading file: ", err.Error())
				os.Exit(1)
			}
		}
		if request.method == "POST" {
			err := os.WriteFile(path.Join(*dir, filename), request.body, 0644)
			if err != nil {
				fmt.Println("Error writing file: ", err.Error())
				os.Exit(1)
			}
			FileCreatedResponse().send(conn)
		}
	default:
		NotFoundResponse().send(conn)
	}
	conn.Close()
}
