package main

import (
	"bufio"
	"fmt"
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

func GetBingoes(c0, c1 string) (int, error) {
	winningNumbers, err := ParseNumberCard(c0)
	if err != nil {
		return 0, err
	}

	myNumbers, err := ParseNumberCard(c1)
	if err != nil {
		return 0, err
	}

	bingoes := 0
	for _, winningNumber := range winningNumbers {
		if _, bingo := slices.BinarySearch(myNumbers, winningNumber); bingo {
			bingoes++
		}
	}

	return bingoes, nil
}

func GetTotalWinningCards(s *bufio.Scanner) (int, error) {
	wonCards := map[int]int{}
	idCard := 0

	for s.Scan() {
		idCard++
		if _, ok := wonCards[idCard]; !ok {
			wonCards[idCard] = 1
		} else {
			wonCards[idCard]++
		}

		g := strings.Split(s.Text(), ": ")
		c := strings.Split(g[1], " | ")
		bingoes, err := GetBingoes(c[0], c[1])
		if err != nil {
			return 0, err
		}

		if bingoes == 0 {
			continue
		}

		for i := 1; i <= bingoes; i++ {
			if _, ok := wonCards[idCard+i]; !ok {
				wonCards[idCard+i] = wonCards[idCard]
				continue
			}

			wonCards[idCard+i] += wonCards[idCard]
		}
	}

	totalWinningCards := 0
	for _, c := range wonCards {
		totalWinningCards += c
	}

	return totalWinningCards, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetTotalWinningCards(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
