package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"

	go func() {
		fmt.Println(msg)
	}()

	time.Sleep(100 * time.Millisecond)
}
