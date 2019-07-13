package jedilevel5

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v1        vehicle
	fourWheel bool
}

type sedan struct {
	v2     vehicle
	luxury bool
}

// StructOperations ... adding a struct data type into a struct
func StructOperations() {

	base := truck{
		v1: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	temp := sedan{
		v2: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(base.v1.doors, "is the number of doors", base.v1.color, "is the color of", base.fourWheel)
	fmt.Println(temp.v2.doors, "is the number of doors", temp.v2.color, "is the color of", temp.luxury)
}
