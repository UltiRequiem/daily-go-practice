package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func getRandomTime() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5)
}

func (c *Consumer) queue(input int) bool {
	c.jobsChan <- input
	fmt.Println("already send input value:", input)
	return true
}

func (c *Consumer) worker(num int) {
	for job := range c.jobsChan {
		n := getRandomTime()
		fmt.Printf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Println("worker:", num, " job value:", job)
	}
}

const poolSize = 2

func main() {
	consumer := Consumer{
		inputChan: make(chan int, 1),
		jobsChan:  make(chan int, poolSize),
	}

	for i := 0; i < poolSize; i++ {
		go consumer.worker(i)
	}

	consumer.queue(1)
	consumer.queue(2)
	consumer.queue(3)
	consumer.queue(4)
	consumer.queue(5)

	time.Sleep(5 * time.Second)
}
