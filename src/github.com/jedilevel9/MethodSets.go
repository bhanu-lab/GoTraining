package jedilevel9

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("This guy can speak a lot")
}

func saySomething(h human) {
	h.speak()
}

func proveReceiverAndValueRelation() {
	p := person{
		"peter",
		"parker",
		30,
	}

	saySomething(&p)
	//if speak function has receiver of type person then saysomething func can receive
	//both pointer and noirmal values but if speak function has pointer person then
	//only pointer types are allowed inside saysomething func
	//saySomething(p)
}
