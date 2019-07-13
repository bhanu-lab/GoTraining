package jedilevel6

import "testing"

func TestCallingBackAFunction(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	expectedValue := 15
	actualValue := CallingBackAFunction(x...)

	if expectedValue != actualValue {
		t.Fail()
	}
}
