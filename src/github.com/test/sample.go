package main

import "fmt"

func main() {
	c := make(chan int)

	go soldier(c)

	for v := range c {
		fmt.Println(v)
	}
}

func soldier(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}
