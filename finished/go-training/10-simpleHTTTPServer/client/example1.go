package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
