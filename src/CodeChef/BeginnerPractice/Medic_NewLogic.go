package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var oddMonths = []int{1, 3, 5, 7, 8, 10, 12}

func main() {
	var dateString = []string{}
	num, _ := strconv.Atoi(readFromInput())
	for i := 0; num > i; i++ {
		dateString = append(dateString, readFromInput())
	}
	for i := 0; len(dateString) > i; i++ {
		val := findNumberOfDays(dateString[i])

		fmt.Println(val)
	}
}

func findNumberOfDays(dat string) int {

	// fmt.Printf("received date is %v\n", dat)
	ymd := strings.Split(dat, ":")
	// fmt.Printf("received date is %v\n", ymd)
	year, _ := strconv.Atoi(ymd[0])
	m, _ := strconv.Atoi(ymd[1])
	datestr := fmt.Sprintf("%s", ymd[2])
	date, _ := strconv.Atoi(datestr)
	// fmt.Printf("year: %v month: %v date: %v", year, m, date)

	if m == 2 {
		if ok := isLeapYear(year); ok {
			return ((29 - date) / 2) + 1
		} else {
			if date%2 == 0 {
				return ((28 - date) / 2) + 15 + 1
			} else {
				return ((28 - date) / 2) + 16 + 1
			}
		}
	} else if ok := in(oddMonths, m); ok {
		// fmt.Println("in 31 day months")
		return ((31 - date) / 2) + 1
	} else {
		// fmt.Println("in 30 day months")
		if date%2 == 0 {
			return ((30 - date) / 2) + 15 + 1
		} else {
			return ((30 - date) / 2) + 16 + 1
		}
	}
}

// readFromInput ... reads input from console
func readFromInput() string {
	rdr := bufio.NewReader(os.Stdin)
	text, _ := rdr.ReadString('\n')
	return strings.Trim(text, "\n")
	//return text
}

//isLeapYear... takes input ass year and returns a bool which tells whether given
//year is Leap year or not
func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		// fmt.Println("its is a leap year")
		return true
	} else if year%100 == 0 && year%400 == 0 {
		// fmt.Println("its is a leap year")
		return true
	} else {
		// fmt.Println("its not a leap year")
		return false
	}
}

//in ... function to check if the val is present in the list sent
func in(list interface{}, val interface{}) bool {
	// fmt.Println("month value is ", val)
	// fmt.Println("list value is ", list)
	dataType := fmt.Sprintf("[]%T", val)
	fmt.Println(dataType)
	listVal := reflect.ValueOf(list)
	actualList := listVal.Interface().([]int)
	for i := 0; i < len(actualList); i++ {
		// fmt.Printf("list value is %v - %Tand value is %v - %T\n", actualList[i], actualList[i], val, val)
		if actualList[i] == val {
			return true
		}
	}
	return false
}
