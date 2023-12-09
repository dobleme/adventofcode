package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	input    []string
	expected int
}{
	{[]string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}, 2},
}

func TestGetSumOfPreviousInSequences(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			res, err := GetSumOfPreviousInSequences(tt.input)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}
var res int

func BenchmarkGetSumOfPreviousInSequences(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r, err := GetSumOfPreviousInSequences(benchCase)
		if err != nil {
			b.Error(err)
		}
		res = r
	}
}
