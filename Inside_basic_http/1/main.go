package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//read request
	request(conn)
	//read response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		res := scanner.Text()
		fmt.Println(res)
		if i == 0 {
			sl := strings.Fields(res)[0]
			fmt.Println("METHOD__", sl)
		}
		if res == " " {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello world</title>
	</head>
	<body>
	<h1>Krish</h1>
	</body>
	</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
