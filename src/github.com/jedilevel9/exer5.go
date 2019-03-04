package jedilevel9

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count1 int64
var waitress1 sync.WaitGroup

func RemoveRaceConditionusingAtomicity() {
	waitress.Add(start)

	for i := 0; i < start; i++ {
		go func() {
			num := atomic.LoadInt64(&count1)
			num = atomic.AddInt64(&count1, 1)

			atomic.StoreInt64(&count1, num)
			waitress1.Done()
		}()
	}

	waitress1.Wait()
	fmt.Println(counter)
}
