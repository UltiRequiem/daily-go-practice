package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.NumCPU() is a call expression
	// console.log(os.cpus().length) On Nodejs
	fmt.Println(runtime.NumCPU() + 1)
}
