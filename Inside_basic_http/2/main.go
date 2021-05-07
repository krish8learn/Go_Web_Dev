package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == " " {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0] //method
	u := strings.Fields(ln)[1] //uri
	fmt.Println("****METHOD", m)
	fmt.Println("****URI", u)

	if m == "GET" && u == "/" {
		startpage(conn)
	} else if m == "GET" && u == "/about" {
		about(conn)
	} else if m == "POST" && u == "/myself" {
		myself(conn)
	}
}

func startpage(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title></title>
	</head>
	<body>
	<strong>STARTPAGE</strong><br>
	<a href= "/">startpage</a><br>
	<a href= "/about">about</a><br>
	<a href= "/myself">myself</a><br>
	</body>
	</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title></title>
	</head>
	<body>
	<strong>ABOUT</strong><br>
	<a href= "/about">about</a><br>
	<a href= "/">startpage</a><br>
	<a href= "/myself">myself</a><br>
	<form method="post" action="/myself">
	<input type="submit" value="myself">
	</form>
	</body>
	</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func myself(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title></title>
	</head>
	<body>
	<strong>MYSELF</strong><br>
	<a href= "/about">about</a><br>
	<a href= "/myself">myself</a><br>
	<form method="post" action="/myself">
	<input type="submit" value="myself">
	</form>
	</body>
	</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
