package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "\n This is the server,conversation starts\n")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		res := scanner.Text()
		fmt.Println(res)
		fmt.Fprintf(conn, "You have written:--> %s\n", res)
	}
	defer conn.Close()
}
