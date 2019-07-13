package jedilevel8

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSONEncode ... uses JSON -> NewEncoder which return *Encoder -> then invokes Encode({} interface)
func JSONEncode() {

	u1 := user{
		First: "Peter",
		Last:  "Parker",
		Age:   32,
	}

	u2 := user{
		First: "cabey",
		Last:  "deysey",
		Age:   30,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println(err)
	}

}
