package jedilevel9

import (
	"fmt"
	"sync"
)

var count int
var waitress sync.WaitGroup
var start = 100
var mt sync.Mutex

func RemoveRaceCondition() {
	waitress.Add(start)

	for i := 0; i < start; i++ {
		go func() {
			mt.Lock()
			num := count
			num++
			count = num
			mt.Unlock()
			waitress.Done()
		}()
	}

	waitress.Wait()
	fmt.Println(counter)
}
