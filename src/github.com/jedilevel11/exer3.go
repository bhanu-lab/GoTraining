package jedilevel11

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return ce.info
}

func CustomError() {

}

func foo(e error) {
	fmt.Println("foo ran ", e)
}
