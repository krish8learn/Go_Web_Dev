package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	select {
	case <-time.After(200 * time.Millisecond):
		//fmt.Println("hello")
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusBadRequest)
	}

	//fmt.Fprintln(w, `hello`)
}
