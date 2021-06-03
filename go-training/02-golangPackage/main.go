package main

import (
	"fmt"

	"github.com/UltiRequiem/daily-go-practice/go-training/02-golangPackage/controller"
)

func main() {
	fmt.Println("Greetings")

	hi := controller.HelloWorld("Zero")
	fmt.Println(hi)
}
