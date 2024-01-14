package main

import (
	"fmt"
	"net"
	"strconv"
)

type Response struct {
	version    string
	statusCode int
	status     string
	headers    map[string]string
	body       string
}

func EmptyOkResponse() *Response {
	return OkResponse("")
}

func OkResponse(body string) *Response {
	headers := make(map[string]string)
	if body != "" {
		headers["Content-Type"] = "text/plain"
		headers["Content-Length"] = strconv.Itoa(len(body))
	}
	return &Response{
		version:    "HTTP/1.1",
		statusCode: 200,
		status:     "OK",
		headers:    headers,
		body:       body,
	}
}

func (r *Response) send(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("%s %s %s\r\n", r.version, strconv.Itoa(r.statusCode), r.status)))
	for header, value := range r.headers {
		conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", header, value)))
	}
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(r.body))
}
