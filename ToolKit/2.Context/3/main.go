package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-time.After(2* time.Second):
		fmt.Println("hello")
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	}
}

//WithTimeout always call with deadline inside           
