package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 21
		close(ch)
	}()

	if v, ok := <-ch; ok == true {
		fmt.Println("value is ", v, ok)
	}

	if v, ok := <-ch; ok == false {
		fmt.Println("value is ", v, ok)
	}
}
