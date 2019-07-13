package jedilevel8

import "testing"

func TestSortUsers(t *testing.T) {
	SortUsers()
	t.Fail() //intentionally failling to check stdout
}
