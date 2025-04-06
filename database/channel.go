package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 5) //can add a second argument for buffer size - process func can finish early
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c) //do this stuff right before the function exits (like closing a file reader)
	for i := 0; i < 5; i++ {
		c <- i
	}
}