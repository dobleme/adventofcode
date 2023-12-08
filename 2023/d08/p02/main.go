package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type tuple struct {
	left  string
	right string
}

func findTotalSteps(current string, route []rune, m map[string]tuple) int {
	totalSteps := 0
	step := 0
	for {
		if step%len(route) == 0 {
			step = 0
		}

		switch route[step] {
		case 'L':
			current = m[current].left
		case 'R':
			current = m[current].right
		}

		totalSteps++
		step++

		if strings.HasSuffix(current, "Z") {
			break
		}
	}

	return totalSteps
}

// Copy Paste from https://go.dev/play/p/SmzvkDjYlb
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Copy Paste from https://go.dev/play/p/SmzvkDjYlb
func LCM(a, b int, c ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(c); i++ {
		result = LCM(result, c[i])
	}

	return result
}

func GetTotalSteps(s []string) (int, error) {
	route := []rune(s[0])
	m := make(map[string]tuple, len(s)-2)
	currents := []string{}
	for i := 2; i < len(s); i++ {
		key := s[i][:3]
		m[key] = tuple{
			left:  s[i][7:10],
			right: s[i][12:15],
		}

		if strings.HasSuffix(key, "A") {
			currents = append(currents, key)
		}
	}

	sols := make([]int, len(currents))
	for i, c := range currents {
		sols[i] = findTotalSteps(c, route, m)
	}

	totalSteps := LCM(sols[0], sols[1], sols[2:]...)
	return totalSteps, nil
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

	res, err := GetTotalSteps(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
