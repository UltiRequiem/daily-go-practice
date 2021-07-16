package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi, golang")

	res, err := http.Get("https://github.com/UltiRequiem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
