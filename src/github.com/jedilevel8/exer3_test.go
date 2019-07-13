package jedilevel8

import (
	"testing"
)

func TestJSCONEncoder(t *testing.T) {
	JSONEncode()
	t.Fail() // Failing puposefully to check out std out
}
