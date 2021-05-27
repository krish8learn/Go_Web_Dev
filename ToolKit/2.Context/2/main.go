package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("hello")
	case <-ctx.Done():
		log.Fatalln(ctx.Err().Error())
	}
	/*
		//below function print after 2 seconds
		time.AfterFunc(2*time.Second, func() {
			fmt.Println("hello")
		})*/
}
