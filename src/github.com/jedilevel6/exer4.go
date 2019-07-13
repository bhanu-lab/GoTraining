package jedilevel6

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// ReceiverOnFunction ... receiver defined ona function and attached to struct type
func ReceiverOnFunction() {

	peter := person{
		first: "peter",
		last:  "parker",
		age:   32,
	}

	peter.speak()

}

// func (r receiver) identifier(params...) returnType {}
func (p person) speak() {
	fmt.Println("Hi I am ", p.first, p.last, "and my age is", p.age)
}
