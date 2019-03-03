package jedilevel9

import (
	"testing"
)

func TestPrintSystemDetails(t *testing.T) {
	PrintSystemDetails()
	t.Fail() // failing deliberately to print println statements
}
