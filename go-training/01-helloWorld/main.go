package main

import "fmt"

func main() {
	mt.Printf(helloWorld("Zero Requiem!"))
	fmt.Println("This is Golang.")
	a := 1
	if a >= 1 {
		fmt.Println("a >= 1")
	}
}

func helloWorld(userName string) string {
	return fmt.Sprintf("Hi, %s ", userName)
}
