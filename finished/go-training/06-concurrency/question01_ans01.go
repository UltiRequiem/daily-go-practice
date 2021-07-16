package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Let's Go")
	}()

	time.Sleep(1 * time.Second)
}
