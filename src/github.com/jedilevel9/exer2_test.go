package jedilevel9

import (
	"testing"
)

func TestPrintConcurrently(t *testing.T) {
	PrintConcurrently()
	t.Fail()
}
