package main

/*
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	testcases := readFromInput()
	testcases = strings.Trim(testcases, "\n")
	test, _ := strconv.Atoi(strings.Trim(testcases, "\n"))
	// fmt.Printf("number of test %v\n", test)
	checkNumberOfCorrectDays(test)
}

// readFromInput ... reads input from console
func readFromInput() string {
	rdr := bufio.NewReader(os.Stdin)
	text, _ := rdr.ReadString('\n')
	return strings.Trim(text, "\n")
}

//checkNumberOfCorrectDays ...
func checkNumberOfCorrectDays(test int) {
	iterations := make([]string, test)
	for index := 0; test > index; index++ {
		iterations[index] = readFromInput()
	}
	// fmt.Println(iterations, len(iterations))
	for z := 0; len(iterations) > z; z++ {
		// fmt.Printf("after reading dates %v and z value is %v\n", iterations[z], z)
		dates := strings.Split(iterations[z], ":")
		// fmt.Println(dates)
		year, _ := strconv.Atoi(dates[0])
		month, _ := strconv.Atoi(dates[1])
		date, _ := strconv.Atoi(dates[2])

		// actual start date of medicine intake
		startDate := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
		var counter, nextDate int
		counter = 0
		previousDate := date

		//loops until current date and alternate date are both odd
		for {
			afterAdding := startDate.AddDate(0, 0, 1)
			_, _, nextDate = afterAdding.Date()
			if previousDate%2 == 1 {
				counter++
			}

			if previousDate%2 == 1 && nextDate%2 == 1 {
				break
			}
			startDate = afterAdding
			previousDate = nextDate
		}
		fmt.Println(counter)
	}
}
*/
