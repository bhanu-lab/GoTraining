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

func StructOperations() {

	base := truck{
		vehicle{
			doors: 2,
			color: "black",
		},
		true,
	}

	temp := sedan{
		vehicle{
			doors: 4,
			color: "red",
		},
		true,
	}

	fmt.Println(base.doors, "is the number of doors", base.color, "is the color of", base.fourWheel)
	fmt.Println(temp.doors, "is the number of doors", temp.color, "is the color of", temp.luxury)
}
