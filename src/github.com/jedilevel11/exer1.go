package jedilevel11

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func CheckForError() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Println("Exception occured while marshalling")
		return
	}
	fmt.Println(string(bs))
}
