package main

import "fmt"

func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO")
		<-c
	}()
	c <- true
}
