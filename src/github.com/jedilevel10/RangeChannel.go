package jedilevel10

import "fmt"

func RangeChannels() {

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
}
