package dog

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns total number of words present in a string separated by space
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	/* m := make(map[string]int)
	for _, v := range xs {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		}
	} */
	return len(xs)
}
