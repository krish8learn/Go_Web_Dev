package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID",007)
	ctx = context.WithValue(ctx,"fname", "james")

	res:= access(ctx)
	fmt.Fprintln(w, res)
}

func access(ctx context.Context)(int) {
	//taking the Value and converting into int
	uid:= ctx.Value("userID").(int)
	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println(ctx)
	fmt.Fprintln(w, ctx)
}
