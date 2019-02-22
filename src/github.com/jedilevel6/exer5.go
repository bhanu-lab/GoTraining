package jedilevel6

import (
	"fmt"
	"math"
	"strconv"
)

type square struct {
	side float64
}

// struct type CIRCLE
type CIRCLE struct {
	radius float64
}

func (s square) area() float64 {
	fmt.Println("area of the square is ", s.side*s.side)
	return float64(s.side * s.side)
}

func (c CIRCLE) area() float64 {
	num, _ := strconv.ParseFloat((fmt.Sprintf("%.2f", math.Pi*math.Pow(c.radius, 2))), 64)
	fmt.Println("area of the circle is ", num)
	return num
}

// SHAPE interface for polymorphism
type SHAPE interface {
	area() float64
}

// AreaBasedOnShape ... function used to calculated area based on the shape type
func AreaBasedOnShape(shape SHAPE) float64 {

	area := 0.1
	switch shape.(type) {
	case square:
		//area = (shape.(square).side)
		area = shape.(square).area()
	case CIRCLE:
		//area = (shape.(CIRCLE).radius)
		area = (shape.(CIRCLE).area())
	}

	return area
}
