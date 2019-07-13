package jedilevel10

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func testRoutinesIncAndDec() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(2)

	wg.Add(2)

	go func() {
		for i := 1; i <= 1000; i++ {
			fmt.Println("INC Vakue is ", i)
			time.Sleep(2 * time.Microsecond)
		}
		fmt.Println("Incrementer is completed")
		wg.Done()
	}()

	go func() {
		for i := 1000; i >= 1; i-- {
			fmt.Println("DEC Value is ", i)
			time.Sleep(2 * time.Microsecond)
		}
		fmt.Println("Decrementer is completed")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Incrementer and Decrementer is completed")
}
