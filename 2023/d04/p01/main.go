package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func ParseNumberCard(s string) ([]int, error) {
	s = strings.ReplaceAll(strings.TrimSpace(s), "  ", " ")
	numbers := strings.Split(s, " ")
	numbersCard := make([]int, len(numbers))

	for i, num := range numbers {
		n, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return nil, err
		}

		numbersCard[i] = n
	}

	slices.Sort(numbersCard)

	return numbersCard, nil
}

func GetWinningCardsPoints(s *bufio.Scanner) (int, error) {
	totalPoints := 0
	for s.Scan() {
		g := strings.Split(s.Text(), ": ")
		c := strings.Split(g[1], " | ")

		winningNumbers, err := ParseNumberCard(c[0])
		if err != nil {
			return 0, err
		}

		myNumbers, err := ParseNumberCard(c[1])
		if err != nil {
			return 0, err
		}

		bingoes := 0
		for _, winningNumber := range winningNumbers {
			if _, bingo := slices.BinarySearch(myNumbers, winningNumber); bingo {
				bingoes++
			}
		}

		if bingoes > 0 {
			totalPoints += int(math.Pow(2, float64(bingoes-1)))
		}
	}

	return totalPoints, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetWinningCardsPoints(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
