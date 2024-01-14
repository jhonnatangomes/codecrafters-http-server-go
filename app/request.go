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
	for line, err := reader.ReadString('\n'); line != "\r\n"; line, err = reader.ReadString('\n') {
		if err != nil {
			fmt.Println("Error reading header: ", err.Error())
			os.Exit(1)
		}
		parts := strings.Split(line, ": ")
		headers[parts[0]] = strings.Trim(parts[1], "\r\n")
	}
	return &Request{
		method:  parts[0],
		path:    parts[1],
		version: parts[2],
		headers: headers,
	}
}
