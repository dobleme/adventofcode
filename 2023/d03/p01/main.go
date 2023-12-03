package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsSymbol(r rune) bool {
	return !IsNumber(r) && r != '.'
}

func isAdjacentNumber(row int, startPos int, endPos int, sc [][]rune) bool {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i == len(sc) {
			continue
		}

		for j := startPos - 1; j <= endPos+1; j++ {
			if j < 0 || j == len(sc[i]) {
				continue
			}

			if IsSymbol(sc[i][j]) {
				return true
			}
		}

	}

	return false
}

func GetSumAdjacentNumbers(s *bufio.Scanner) (int, error) {
	var sc [][]rune
	for s.Scan() {
		sc = append(sc, []rune(s.Text()+"."))
	}

	sumAdjacentNumber := 0
	for i, row := range sc {
		number := ""
		startPos := 0

		for j, r := range row {
			if IsNumber(r) {
				if number == "" {
					startPos = j
				}

				number += string(r)
				continue
			}

			if number == "" {
				continue
			}

			if isAdjacentNumber(i, startPos, j-1, sc) {
				n, err := strconv.Atoi(number)
				if err != nil {
					return 0, err
				}

				sumAdjacentNumber += n
			}

			number = ""
			startPos = 0
		}
	}

	return sumAdjacentNumber, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetSumAdjacentNumbers(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
