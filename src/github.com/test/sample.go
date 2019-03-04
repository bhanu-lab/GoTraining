package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count1 int64
var waitress1 sync.WaitGroup
var start = 100

func main() {
	waitress1.Add(start)

	for i := 0; i < start; i++ {
		go func() {
			atomic.AddInt64(&count1, 1)
			waitress1.Done()
		}()
	}

	waitress1.Wait()
	fmt.Println(count1)
}
