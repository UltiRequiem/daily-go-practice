package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Let's Go")
		c <- true
	}()

	<-c
}
