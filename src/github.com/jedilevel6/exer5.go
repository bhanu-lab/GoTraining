package jedilevel6

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type CIRCLE struct {
	radius float64
}

func (s square) area() float64 {
	fmt.Println("area of the square is ", s.side*s.side)
	return float64(s.side * s.side)
}

func (c CIRCLE) area() float64 {
	fmt.Println("area os the square is ", (math.Pi * math.Pow(c.radius, 2)))
	return math.Pi * math.Pow(c.radius, 2)
}

type SHAPE interface {
	area() float64
}

func AreaBasedOnShape(shape SHAPE) float64 {

	area := 0.1
	switch shape.(type) {
	case square:
		area = (shape.(square).area)
	case CIRCLE:
		area = (shape.area)
	}

	return area
}
