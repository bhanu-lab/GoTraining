package jedilevel10

import "fmt"

func CheckChannelsUsage() {

	ch := make(chan int)

	//channels can be converter from generic to specific
	go sender(ch)

	//not running in go routine to make sure we receive the value once channel is updated with value
	receiver(ch)

	fmt.Println("Done with execution")
}

func sender(ch chan<- int) {
	ch <- 99
}

func receiver(ch <-chan int) {
	fmt.Println(<-ch)
}
