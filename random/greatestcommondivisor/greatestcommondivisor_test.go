package greatestcommondivisor_test

import (
	"github.com/meshachdamilare/dsapratice/random/greatestcommondivisor"
	"testing"
)

func TestGCD(t *testing.T) {
	tt := []struct {
		a        int
		b        int
		expected int
	}{
		{4, 12, 4},
		{50, 10, 10},
	}

	t.Log("Given the need to test greatest common divisor")

	for testId, test := range tt {
		got := greatestcommondivisor.GCD(test.a, test.b)
		if got != test.expected {
			t.Fatalf("TestID: %d\tExpected:  %d Got: %d\n", testId, test.expected, got)
		}
		t.Logf("TestID: %d\tExpected: %d Got: %d", testId, test.expected, got)
	}
}
