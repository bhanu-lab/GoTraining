package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func TestYears(t *testing.T) {
	v := Years(10)

	if v != 70 {
		t.Fatal("Expected is ", 70, " output is ", v)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

// go test -cover gives total percent of  coverage of statements
// go test -coverprofile c.out writes coverage profile to c out which can be used to generate html report
// go tool cover -html=c.out provides coverage map on browser

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(10)

	if v != 70 {
		t.Fatal("Expected is ", 70, " output is ", v)
	}
}
