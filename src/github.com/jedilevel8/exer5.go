package jedilevel8

import (
	"fmt"
	"sort"
)

type person struct {
	First   string
	Last    string 
	Age     int
	Sayings []string
}

type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type byLast []person

func (b byLast) Len() int {
	return len(b)
}

func (b byLast) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLast) Less(i, j int) bool {
	return b[i].Last < b[j].Last
}

// CustomSort ... implemts custom sort for last name and age
func CustomSort() {
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	for _, v := range users {
		sort.Strings(person(v).Sayings)
	}

	fmt.Println("Before Sort by Age\n", users)

	sort.Sort(byAge(users))

	fmt.Println("After Sort Age\n", users)

	fmt.Println("")

	fmt.Println("Before Sort by Last\n", users)

	sort.Sort(byLast(users))

	fmt.Println("After Sort Last\n", users)

}
