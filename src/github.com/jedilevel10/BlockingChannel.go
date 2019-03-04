package jedilevel10

import "fmt"

func checkChannels() {
	ch := make(chan int)

	ch <- 42

	fmt.Println(<-ch)
}
