package main

import (
	"fmt"

	"github.com/UltiRequiem/daily-go-practice/11-crossBuild/hello"
	"github.com/UltiRequiem/daily-go-practice/11-crossBuild/hello2"
	"github.com/UltiRequiem/daily-go-practice/11-crossBuild/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
