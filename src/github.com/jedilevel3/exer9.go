package jedilevel3

import (
	"fmt"
)

// DetermineFavSportUsingSwitchStatement ... prints
func DetermineFavSportUsingSwitchStatement(favsport string) {

	switch favsport {

	case "cricket":
		fmt.Println("My favourite sport is cricket")

	case "hocket":
		fmt.Println("My favourite sport is hocket")

	case "badminton":
		fmt.Println("My favourite sport is badminton")

	case "football":
		fmt.Println("My favourite sport is football")

	default:
		fmt.Println("I actually dont care")

	}

}
