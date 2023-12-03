package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

var testCases = []struct {
	input    string
	expected int
}{
	{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`, 467835},
}

func TestEngineSchemaGearRatios(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			s := bufio.NewScanner(strings.NewReader(tt.input))
			res, err := GetSumGearRatios(s)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func BenchmarkEngineSchemaGearRatios(b *testing.B) {
	s := bufio.NewScanner(strings.NewReader(benchCase))
	for n := 0; n < b.N; n++ {
		_, err := GetSumGearRatios(s)
		if err != nil {
			b.Error(err)
		}
	}
}
