package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	i := 0
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1 done")
		i++
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2 done")
		i++
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 3 done")
		i++
	}()
	wg.Wait()
	fmt.Println("all goroutine done")
	fmt.Println(i)
}
