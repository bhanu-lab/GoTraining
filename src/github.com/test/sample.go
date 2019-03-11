package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/* ch := make(chan int)

	for i := 0; i < 10; i++ { // creates 10 go routines and adds to waitgroup
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- j * i
			}
		}(i)
	}

	fmt.Println(runtime.NumGoroutine())

	for v := 0; v < 100; v++ {
		fmt.Println(<-ch)
	}
	fmt.Println("About to exit program ...") */

	ch := make(chan int)

	for i := 0; i < 10; i++ { // creates 10 go routines and adds to waitgroup
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				ch <- j
			}
			wg.Done() // indication of go routine is done to main routine
		}()
	}

	fmt.Println(runtime.NumGoroutine())
	go func() {
		wg.Wait() //wait for all go routines to complete
		close(ch) // closing channel after completion of wait fo go routines
	}()
	for v := range ch { // range can be used since channel is closed
		fmt.Println(v)
	}
	fmt.Println("About to exit program ...")

}
