package jedilevel9

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var waiter sync.WaitGroup
var invalue = 100

func main() {
	waiter.Add(invalue)

	for i := 0; i < invalue; i++ {
		go func() {
			num := counter
			runtime.Gosched()
			num++
			counter = num
			waiter.Done()
		}()
	}

	waiter.Wait()
	fmt.Println(counter)
}
