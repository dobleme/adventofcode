package main

import (
	"bufio"
	"fmt"
	"os"
)

type tuple struct {
	left  string
	right string
}

func GetTotalSteps(s []string) (int, error) {
	route := []rune(s[0])
	m := make(map[string]tuple, len(s)-2)
	for i := 2; i < len(s); i++ {
		m[s[i][:3]] = tuple{
			left:  s[i][7:10],
			right: s[i][12:15],
		}
	}

	totalSteps := 0
	step := 0
	current := "AAA"
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

		if current == "ZZZ" {
			break
		}
	}

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
