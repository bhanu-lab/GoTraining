package jedilevel8

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Last  string
	Age   int
}

func MarshalToUser() {

	u1 := user{
		First: "Peter",
		Last:  "Parker",
		Age:   32,
	}

	u2 := user{
		First: "Sharon",
		Last:  "Tuchey",
		Age:   30,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	bs, error := json.Marshal(users)

	if error != nil {
		fmt.Println("error occured", error)
	}

	fmt.Println(string(bs))
}
