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
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}, 2},
	{[]string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}, 6},
}

func TestGetTotalSteps(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			res, err := GetTotalSteps(tt.input)
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
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}
var res int

func BenchmarkGetTotalSteps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r, err := GetTotalSteps(benchCase)
		if err != nil {
			b.Error(err)
		}
		res = r
	}
}
