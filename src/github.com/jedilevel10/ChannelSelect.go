package jedilevel10

import "fmt"

// ChannelSelection ...
func ChannelSelection() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//transmitter
	go transmitter(eve, odd, quit)

	consumer(eve, odd, quit)

	fmt.Println("Exiting the Program")
}

func transmitter(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
}

func consumer(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from even channel is ", v)
		case v := <-odd:
			fmt.Println("the value received from odd channel is ", v)
		case v := <-quit:
			fmt.Println("value received from quit channel is ", v)
			return
		}
	}
}
