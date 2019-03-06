package jedilevel10

import "fmt"

func OkIdiom() {
	ch := make(chan int)

	go func() {
		ch <- 21
	}()

	if v, ok := <-ch; ok == true {
		fmt.Println("value is ", v, ok)
	}
}
