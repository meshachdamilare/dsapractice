package containsduplicate_test

import (
	"github.com/meshachdamilare/dsapratice/random/containsduplicate"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"Has", []int{1, 2, 3, 3, 4}, true},
		{"Not", []int{1, 2, 3, 4, 5}, false},
	}
	t.Log("Given the need to test if a slice contain Duplicate")
	{
		for testID, test := range tt {
			t.Logf("\tTest %d:\tWhen checking the value %v.", testID, test.input)
			got := containsduplicate.ContainsDuplicates(test.input)
			if got != test.expected {
				t.Logf("\t%v\tTest %d:\tShould be false if input has no duplicates", test.expected, testID)
				t.Fatalf("\t\tTest %d:\tGot %v, Excpeted %v", testID, got, test.expected)
			}
			t.Logf("\t%v\tTest %d:\tShould be false if the input has duplicates", test.expected, testID)
		}
	}
}
