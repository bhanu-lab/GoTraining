package jedilevel11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddTwoNumbers() {

	fmt.Println("Enter first number: ")
	num1 := checkForValidInteger()
	fmt.Println("Enter second number: ")
	num2 := checkForValidInteger()
	total := num1 + num2

	fmt.Println("sum of two numbers is: ", total)

}

func readAnInteger() (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	num1, _ := reader.ReadString('\n')
	num1 = strings.Trim(num1, "\n")
	x, err := strconv.ParseInt(num1, 10, 64)

	if err != nil {
		return 0, err
	}

	return x, nil
}

func checkForValidInteger() int64 {

	for {
		num1, err := readAnInteger()
		if err != nil {
			fmt.Println("An error occured", err)
			continue
		}
		return num1
	}

}

func readNumber() int64 {
	var val string
	_, err := fmt.Scanln(&val)
	if err != nil {
		fmt.Println("Error occured while trying to scan value")
		return 0
	}
	fmt.Println("value read is")
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
