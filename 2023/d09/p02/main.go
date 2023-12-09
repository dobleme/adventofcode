package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseSeqs(s []string) ([][]int, error) {
	seqs := make([][]int, len(s))
	for i, l := range s {
		seq := strings.Split(l, " ")
		seqs[i] = make([]int, len(seq))

		for j, n := range seq {
			number, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}

			seqs[i][j] = number

		}
	}

	return seqs, nil
}

func differences(seq []int) []int {
	diffs := make([]int, len(seq)-1)
	for i := 0; i < len(diffs); i++ {
		diffs[i] = seq[i+1] - seq[i]
	}

	return diffs
}

func allZeroes(seq []int) bool {
	for _, s := range seq {
		if s != 0 {
			return false
		}
	}

	return true
}

func previousNumberOfSeq(seq []int) int {
	diffs := differences(seq)
	if allZeroes(diffs) {
		return seq[0]
	}

	return seq[0] - previousNumberOfSeq(diffs)
}

func GetSumOfPreviousInSequences(s []string) (int, error) {
	seqs, err := parseSeqs(s)
	if err != nil {
		return 0, err
	}

	totalNexts := 0
	for _, seq := range seqs {
		totalNexts += previousNumberOfSeq(seq)
	}

	return totalNexts, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	res, err := GetSumOfPreviousInSequences(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
