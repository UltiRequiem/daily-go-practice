package main

import "fmt"

func main() {
	foo := func() string {
		return "Hello World One."
	}

	fmt.Println(foo())

	bar := func() {
		fmt.Println("Hello World Two.")
	}

	bar()

	func() {
		fmt.Println("Hello World Three.")
	}()

	go func(i, j int) {
		fmt.Println(i + j)
	}(1, 2)
}
