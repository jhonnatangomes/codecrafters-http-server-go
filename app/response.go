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
	body       []byte
}

func EmptyOkResponse() *Response {
	return OkResponse(nil)
}

func FileOkResponse(file []byte) *Response {
	return &Response{
		version:    "HTTP/1.1",
		statusCode: 200,
		status:     "OK",
		headers: map[string]string{
			"Content-Type":   "application/octet-stream",
			"Content-Length": strconv.Itoa(len(file)),
		},
		body: file,
	}
}

func FileCreatedResponse() *Response {
	return &Response{
		version:    "HTTP/1.1",
		statusCode: 201,
		status:     "Created",
		headers:    nil,
		body:       nil,
	}
}

func OkResponse(body []byte) *Response {
	headers := make(map[string]string)
	if body != nil {
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

func NotFoundResponse() *Response {
	return &Response{
		version:    "HTTP/1.1",
		statusCode: 404,
		status:     "Not Found",
		headers:    nil,
		body:       nil,
	}
}

func (r *Response) send(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("%s %s %s\r\n", r.version, strconv.Itoa(r.statusCode), r.status)))
	for header, value := range r.headers {
		conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", header, value)))
	}
	conn.Write([]byte("\r\n"))
	conn.Write(r.body)
}
