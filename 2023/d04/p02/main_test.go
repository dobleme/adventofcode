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
	{`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`, 30},
}

func TestScratchcardsWinningCards(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			s := bufio.NewScanner(strings.NewReader(tt.input))
			res, err := GetTotalWinningCards(s)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func BenchmarkScratchcardsWinningCards(b *testing.B) {
	s := bufio.NewScanner(strings.NewReader(benchCase))
	for n := 0; n < b.N; n++ {
		_, err := GetTotalWinningCards(s)
		if err != nil {
			b.Error(err)
		}
	}
}
