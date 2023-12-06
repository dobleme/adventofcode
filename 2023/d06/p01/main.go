package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func parseNumbers(s string) ([]int, error) {
	var nums []int

	numStr := ""
	for _, r := range s {
		if r >= '0' && r <= '9' {
			numStr += string(r)
		}

		if r == ' ' && numStr != "" {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}

			nums = append(nums, n)
			numStr = ""
		}
	}

	n, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, err
	}

	nums = append(nums, n)

	return nums, nil
}

func GetWaysOfBeatingTheRecord(s []string) (int, error) {
	times, err := parseNumbers(s[0])
	if err != nil {
		return 0, nil
	}

	distance, err := parseNumbers(s[1])
	if err != nil {
		return 0, nil
	}

	totalWays := 1
	for i, t := range times {
		d := distance[i]
		r := math.Sqrt(math.Pow(float64(t), 2) - (float64(4 * -1 * -d)))
		lSol := ((float64(-t) + r) / -2) + 1
		rSol := (float64(-t) - r) / -2

		if rSol == math.Trunc(rSol) {
			rSol -= 1
		}

		totalWays *= int(rSol) - int(lSol) + 1
	}

	return totalWays, nil
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
	res, err := GetWaysOfBeatingTheRecord(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
