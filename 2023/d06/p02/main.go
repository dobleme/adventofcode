package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func parseNumber(s string) (int, error) {
	numStr := ""
	for _, r := range s {
		if r >= '0' && r <= '9' {
			numStr += string(r)
		}
	}

	n, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetWaysOfBeatingTheRecord(s []string) (int, error) {
	time, err := parseNumber(s[0])
	if err != nil {
		return 0, nil
	}

	distance, err := parseNumber(s[1])
	if err != nil {
		return 0, nil
	}

	r := math.Sqrt(math.Pow(float64(time), 2) - (float64(4 * -1 * -distance)))
	lSol := ((float64(-time) + r) / -2) + 1
	rSol := (float64(-time) - r) / -2

	if rSol == math.Trunc(rSol) {
		rSol -= 1
	}

	return int(rSol) - int(lSol) + 1, nil
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
