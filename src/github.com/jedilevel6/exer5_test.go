package jedilevel6

import (
	"testing"
)

func TestAreaBasedOnShapeForSquare(t *testing.T) {
	s := square{
		5,
	}

	expectedSquareArea := float64(25)
	actualAquareArea := AreaBasedOnShape(s)

	if actualAquareArea != expectedSquareArea {

		t.Fail()
	}

}

func TestAreaBasedOnShapeForCircle(t *testing.T) {

	c := CIRCLE{
		8,
	}

	expectedCircleArea := float64(201.06)
	actualacCircleArea := AreaBasedOnShape(c)

	if expectedCircleArea != actualacCircleArea {
		t.Fail()
	}
}
