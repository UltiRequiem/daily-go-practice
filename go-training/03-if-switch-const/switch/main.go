package main

import "fmt"

func checkValue(s int) {
	switch s {
	case 2:
		fallthrough
	case 0, 1:
		fmt.Println("check value is", s)
	default:
		fmt.Println("Nope.")
	}
}

func main() {
	checkValue(0)
	checkValue(1)
	checkValue(2)
	checkValue(8)
}
