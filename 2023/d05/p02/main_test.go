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
	{`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, 46},
}

func TestMinLocationSeed(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			s := bufio.NewScanner(strings.NewReader(tt.input))
			res, err := GetMinLocationSeed(s)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func BenchmarkMinLocationSeed(b *testing.B) {
	scanners := make([]*bufio.Scanner, b.N)
	for n := 0; n < b.N; n++ {
		scanners[n] = bufio.NewScanner(strings.NewReader(benchCase))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := GetMinLocationSeed(scanners[n])
		if err != nil {
			b.Error(err)
		}
	}
}
