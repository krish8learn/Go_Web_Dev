package main

import (
	"bufio"
	"fmt"
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
		fmt.Fprintln(conn, "Write some texts")
		scanner := bufio.NewScanner(conn)
		if scanner.Scan() == true {
			fmt.Println("Response From client", scanner.Text())
		}
		conn.Close()
	}
}
