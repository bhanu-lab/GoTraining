package main

import (
	"sync"

	"github.com/jedilevel11"
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
	jedilevel11.CheckForError()

}
