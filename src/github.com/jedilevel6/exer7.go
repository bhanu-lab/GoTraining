package jedilevel6

// AssignAnonymousFunctionToVariable calling anonymous function
func AssignAnonymousFunctionToVariable(x ...int) int {

	f := func(x ...int) int {
		var total int
		for _, v := range x {
			total = total + v
		}

		return total
	}

	return f(x...)

}
