package jedilevel6

func InvokeReturneddFunction(x ...int) {

	f := TestFunction()
	f(x...)

}

func TestFunction() func(x ...int) int {
	f := func(x ...int) int {
		var total int
		for _, v := range x {
			total = total + v
		}

		return total
	}
	return f
}
