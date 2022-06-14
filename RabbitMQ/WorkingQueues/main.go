package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	body := messageFromUser(os.Args)
	fmt.Println(body)

}

func messageFromUser(args []string) string {

	var str string
	if len(args) < 2 || os.Args[1] == "" {
		str = "Default Hello"
	} else {
		str = strings.Join(args[1:], " ")
	}

	return str
}
