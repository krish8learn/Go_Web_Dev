package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	//reading from the server
	//bs, err := ioutil.ReadAll(conn)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(string(bs))
	fmt.Fprintln(conn, "I dialed you")
}
