package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	method  string
	path    string
	version string
	headers map[string]string
	body    []byte
}

func newRequest(conn net.Conn) *Request {
	reader := bufio.NewReader(conn)
	startLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading start line: ", err.Error())
		os.Exit(1)
	}
	parts := strings.Split(startLine, " ")
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading header: ", err.Error())
			os.Exit(1)
		}
		if line == "\r\n" {
			break
		}
		parts := strings.Split(line, ": ")
		headers[parts[0]] = strings.Trim(parts[1], "\r\n")
	}
	body := make([]byte, 0)
	for reader.Buffered() > 0 {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Println("Error reading body: ", err.Error())
			os.Exit(1)
		}
		body = append(body, b)
	}
	return &Request{
		method:  parts[0],
		path:    parts[1],
		version: parts[2],
		headers: headers,
		body:    body,
	}
}
