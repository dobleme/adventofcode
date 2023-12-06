package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	input    []string
	expected int
}{
	{[]string{"Time:      7  15   30", "Distance:  9  40  200"}, 71503},
}

func TestWaysOfBeatingTheRecord(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			res, err := GetWaysOfBeatingTheRecord(tt.input)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = []string{"Time:      7  15   30", "Distance:  9  40  200"}
var res int

func BenchmarkWaysOfBeatingTheRecord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r, err := GetWaysOfBeatingTheRecord(benchCase)
		if err != nil {
			b.Error(err)
		}
		res = r
	}
}
