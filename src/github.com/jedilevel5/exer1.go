package jedilevel5

import "fmt"

type example struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func CreatingAndUsingStruct() {

	x := example{
		firstName:   "Test",
		lastName:    "team",
		favIceCream: []string{"chocolate", "strawberry", "hazzlenut"},
	}

	y := example{
		firstName:   "test1",
		lastName:    "team1",
		favIceCream: []string{"chocolate1", "banana", "custard"},
	}

	fmt.Println(x.firstName)
	fmt.Println(x.lastName)

	for _, v := range x.favIceCream {
		fmt.Println(v)
	}

	fmt.Println(y.firstName)
	fmt.Println(y.lastName)

	for _, v := range y.favIceCream {
		fmt.Println(v)
	}
}
