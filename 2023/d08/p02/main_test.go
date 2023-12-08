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
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
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
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
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
