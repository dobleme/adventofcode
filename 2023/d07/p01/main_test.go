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
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}, 6440},
}

func TestGetTotalWinnigns(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			res, err := GetTotalWinnigns(tt.input)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = []string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}
var res int

func BenchmarkGetTotalWinnigns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r, err := GetTotalWinnigns(benchCase)
		if err != nil {
			b.Error(err)
		}
		res = r
	}
}
