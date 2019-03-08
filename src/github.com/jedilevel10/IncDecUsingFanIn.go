package jedilevel10

import "fmt"

func merger() {
	c := fanInNumbers(Incrementer(), Decrementer())

	for i := range c {
		fmt.Println(i)
	}
}

func Incrementer() <-chan int {
	c := make(chan int)

	for i := 1; i < 100; i++ {
		c <- i
	}

	return c
}

func Decrementer() <-chan int {
	c := make(chan int)

	for i := 100; i >= 1; i++ {
		c <- i
	}

	return c
}

func fanInNumbers(input1, input2 <-chan int) <-chan int {
	c := make(chan int)

	for {
		c <- <-input1
	}

	for {
		c <- <-input2
	}

	return c

}
