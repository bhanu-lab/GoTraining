package jedilevel9

import (
	"fmt"
	"runtime"
)

// PrintSystemDetails ... prints system details using go runtime package
func PrintSystemDetails() {

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("Num of CPU's\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
}
