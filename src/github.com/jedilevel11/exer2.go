package jedilevel11

import (
	"encoding/json"
	"fmt"
)

type human struct {
	First   string
	Last    string
	Sayings []string
}

func ImplementNewErrorMessage() {
	p1 := human{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		fmt.Println("here i am throwing spiked error", err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("This is Spiked Error")
	}

	return bs, err
}
