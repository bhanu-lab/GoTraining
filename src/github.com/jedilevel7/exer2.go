package jedilevel7

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func ChangeStructValue() {

	p := person{
		first: "Peter",
		last:  "Parker",
		age:   32,
	}

	fmt.Println(p)
	fmt.Println(&p)

	ChangeValue(&p)

	fmt.Println(p)

}

func ChangeValue(x *person) {
	var y *person

	y = x
	(*x).first = "Samuel"
	fmt.Println("address of x is ", x)
	fmt.Println("address of y is ", &y)
	y.last = "camel"
	fmt.Println("printing value before")
	fmt.Println(y)
	fmt.Println("printing value after")
}
