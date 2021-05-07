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
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(conn, "Rot 13 Application")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		res := scanner.Text()
		bs := []byte(strings.ToLower(res))
		rotres := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", res, rotres)
	}
}

func rot13(bs []byte) []byte {
	var sl = make([]byte, len(bs))
	for i, v := range bs {
		if v >= 109 {
			sl[i] = v + 13
		} else {
			sl[i] = v - 13
		}
	}
	return sl
}
