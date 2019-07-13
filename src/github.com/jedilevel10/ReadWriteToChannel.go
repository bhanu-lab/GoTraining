package jedilevel10

import (
	"fmt"
)

func WriteAndReadChannel() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for v := range c {
			fmt.Println("value is ", v)
		}
	}()

	fmt.Println("Program about to exit")
}
