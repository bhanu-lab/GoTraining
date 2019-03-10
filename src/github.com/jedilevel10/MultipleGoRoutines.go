package jedilevel10

import (
	"fmt"
	"sync"
)

func CreateMultipleRoutines() {
	ch := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				x := i * j
				ch <- x
			}
			wg.Done()
		}()
	}
	wg.Wait()

	for v := range ch {
		fmt.Println(v)
	}
}
