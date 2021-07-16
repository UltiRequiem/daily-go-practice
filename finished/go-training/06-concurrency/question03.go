package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Let's Go"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Let's GoGoGo"
	time.Sleep(1 * time.Second)
}
