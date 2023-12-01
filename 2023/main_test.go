package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var testCases = []struct {
	text     string
	expected int
}{
	{`1fo2o
kjrqmzv9mmtxhgvsevenhvq7
4gjnmxtrbflgp71
1seven336
8sevengzfvjrhnsb6ddb8ninerkgkxthtfkvbcmqs
rkzlnmzgnk91zckqprrptnthreefourtwo
rkzlnmzgnk9zcfourtwo
`, 439},
	{`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eighthree
`, 364},
	{`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`, 281},
	{`nineight`, 98},
	{`five1oneight`, 58},
	{`sevenine`, 79},
	{`one`, 11},
	{`two`, 22},
	{`three`, 33},
	{`four`, 44},
	{`five`, 55},
	{`six`, 66},
	{`seven`, 77},
	{`eight`, 88},
	{`nine`, 99},
	{`twone`, 21},
	{`eightwo`, 82},
	{`nineight`, 98},
	{`eighthree`, 83},
	{`nineight`, 98},
	{`eeeight`, 88},
	{`oooneeone`, 11},
	{`mtwone8onethreecprdhtdgxvdqcptplmsixtwo`, 22},
	{`one2one`, 11},
	{`sevenhk5oneight`, 78},
	{`ilk7gfive7`, 77},
	{`vj8`, 88},
	{`zxkncsrktnrv74eighttwonine`, 79},
	{`nineeightvhncnbtbp68fone`, 91},
}

func SumNumbersByRegexp(s *bufio.Scanner) (int, error) {
	r := regexp.MustCompile(`\d`)

	total := 0
	for s.Scan() {
		matches := r.FindAllString(s.Text(), -1)
		number := fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1])
		n, err := strconv.Atoi(number)
		if err != nil {
			return total, err
		}
		total += n
	}

	return total, nil
}

func TestLine(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("Test case num #%d", i), func(t *testing.T) {
			s := bufio.NewScanner(strings.NewReader(tt.text))
			res, err := SumNumbersByLine(s)
			if err != nil {
				t.Error(err)
			}

			if res != tt.expected {
				t.Errorf("Expected: %d, but got: %d", tt.expected, res)
			}
		})
	}
}

var benchCase = `1fo2o
kjrqmzv9mmtxhgvsevenhvq7
4gjnmxtrbflgp71
1seven336
8sevengzfvjrhnsb6ddb8ninerkgkxthtfkvbcmqs
rkzlnmzgnk91zckqprrptnthreefourtwo
mtwone8onethreecprdhtdgxvdqcptplmsixtwo
one2one
sevenhk5oneight
ilk7gfive7
vj8
zxkncsrktnrv74eighttwonine
nineeightvhncnbtbp68fone
`

func BenchmarkLine(b *testing.B) {
	s := bufio.NewScanner(strings.NewReader(benchCase))
	for n := 0; n < b.N; n++ {
		_, err := SumNumbersByLine(s)
		if err != nil {
			b.Error(err)
		}
	}
}
