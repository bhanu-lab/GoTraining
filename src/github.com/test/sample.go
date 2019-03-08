package main

import (
	"fmt"
)

func main() {
	//var wg sync.WaitGroup
	//c := make(chan int)

	c := fanInNumbers(Incrementer(), Decrementer())

	for i := 0; i < 199; i++ {
		fmt.Println(<-c)
	}
}

func Incrementer() <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i < 100; i++ {
			c <- i
		}
		fmt.Println("Done with incrementer ...")
		close(c)
	}()
	return c
}

func Decrementer() <-chan int {
	c := make(chan int)

	go func() {
		for i := 100; i >= 1; i-- {
			c <- i
		}
		fmt.Println("Done with decrementer ...")
		close(c)
	}()
	return c
}

func fanInNumbers(input1, input2 <-chan int) <-chan int {
	/* var wg sync.WaitGroup
	wg.Add(2) */
	c := make(chan int)
	go func() {
		for i := range input1 {
			c <- i
		}
		fmt.Println("Done with first fan in ...")
		//wg.Done()
	}()

	go func() {
		for j := range input2 {
			c <- j
		}
		fmt.Println("Done with second fan in ...")
		//wg.Done()
	}()

	/* wg.Wait()
	close(c)
	*/
	return c
}
