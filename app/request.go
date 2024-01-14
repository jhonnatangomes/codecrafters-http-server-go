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
}

func newRequest(conn net.Conn) *Request {
	reader := bufio.NewReader(conn)
	startLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading start line: ", err.Error())
		os.Exit(1)
	}
	parts := strings.Split(startLine, " ")
	return &Request{
		method:  parts[0],
		path:    parts[1],
		version: parts[2],
	}
}
