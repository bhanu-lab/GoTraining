package jedilevel6

import "testing"

func TestAssignAnonymousFunctionToVariable(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedSum := 45
	actualSum := AssignAnonymousFunctionToVariable(x...)
	if expectedSum != actualSum {
		t.Fail()
	}
}
