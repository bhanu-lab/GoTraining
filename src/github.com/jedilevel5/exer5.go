package jedilevel5

import "fmt"

// AnonymousStruct ... creating an anonymous struct and using it
func AnonymousStruct() {

	x := struct {
		age  int
		name string
		fav  string
	}{
		42,
		"James Bond",
		"Bad tastes",
	}

	fmt.Println(x.name, "is ", x.age, "years old and his favourite is ", x.fav)

}
