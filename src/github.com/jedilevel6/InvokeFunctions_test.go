package jedilevel6

import (
	"fmt"
	"testing"
)

func TestInvokeFuntions(t *testing.T) {
	InvokeFuntions()
}

func TestCount(t *testing.T) {
	value := count(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("value is ", value)

}
