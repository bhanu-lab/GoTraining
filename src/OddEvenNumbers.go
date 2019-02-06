package main

import "fmt"

func main() {

	for i:=0; i<=100 ; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			fmt.Println(" is even number")
		}else{
			fmt.Print(i)
			fmt.Println(" is odd number")
		}
	}
}
