package main

import "fmt"

const nice = true

func main() {
	const bad = false
	fmt.Println(nice, bad)
}
