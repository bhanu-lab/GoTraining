package dog

import (
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount(`One Two Two Three Three Three`)

	for k, v := range m {
		switch k {
		case "One":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "Two":
			if v != 2 {
				t.Error("got", v, "expected", 2)
			}
		case "Three":
			if v != 3 {
				t.Error("got", v, "expected", 3)
			}
		}
	}

}
