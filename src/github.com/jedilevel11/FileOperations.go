package jedilevel11

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenFile() {

	//fmt.Println("filename received is", fileName)
	f, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("Error occured while trying to read a file pos 1", err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println("Error occured while trying to read a file pos 2")
		return
	}

	fmt.Println(string(bs))
}
