package main

import "fmt"

type car struct {
	name  string
	color string
}

func (c *car) SetName01(s string) {
	fmt.Printf("SetName01: car address: %p\n", c)
	c.name = s
}

func (c car) SetName02(s string) {
	fmt.Printf("SetName02: car address: %p\n", &c)
	c.name = s
}

func main() {
	toyota := &car{
		name:  "toyota",
		color: "white",
	}

	fmt.Printf("car address: %p\n", toyota)

	fmt.Println(toyota.name)
	toyota.SetName01("foo")
	fmt.Println(toyota.name)
	toyota.SetName02("bar")
	fmt.Println(toyota.name)
	toyota.SetName02("test")
	fmt.Println(toyota.name)
}
