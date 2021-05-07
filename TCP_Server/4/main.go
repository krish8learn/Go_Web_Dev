package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		resp, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(resp, "Write some texts befor 5s")
		go handle(resp)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		reText := scanner.Text()
		fmt.Println(reText)
		fmt.Fprintf(conn, "Receive from you %s\n", reText)
	}
	defer conn.Close()
	fmt.Println("CODE ENDS")
}
