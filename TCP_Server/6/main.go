package main

import (
	"bufio"
	"fmt"
	"io"
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
	io.WriteString(conn, "\nIN-MEMORY DATABASE\n"+"USE:n"+
		"SET key value \n"+
		"GET key \n+"+
		"DEL key \n"+
		"EXAMPLE \n"+
		"SET fav choclate \n"+
		"GET fav \n\n")
	store := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		sltext := strings.Fields(text)
		switch sltext[0] {
		case "SET":
			if len(sltext) != 3 {
				fmt.Fprintf(conn, "WRONG INSTRUCTION")
			}
			store[sltext[1]] = sltext[2]
			fmt.Fprintf(conn, "SAVED\n")
		case "GET":
			cond := false
			for key, value := range store {
				if key == sltext[1] {
					fmt.Fprintf(conn, "RESULT :%s\n", value)
					cond = true
					break
				}
			}
			if cond == false {
				fmt.Fprintf(conn, "RESULT : NOT FOUND\n")
			}
		case "DEL":
			cond := false
			for key, _ := range store {
				if key == sltext[1] {
					delete(store, sltext[1])
					fmt.Fprintf(conn, "DELETED\n")
					cond = true
					break
				}
			}
			if cond == false {
				fmt.Fprintf(conn, "RESULT : NOT FOUND\n")
			}
		default:
			fmt.Fprintf(conn, "INVALID COMMAND")
		}
	}
}
