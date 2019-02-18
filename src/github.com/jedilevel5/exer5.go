package jedilevel5

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

}
