package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

type gearPos struct {
	i int
	j int
}

type gearNumber struct {
	number int
	pos    gearPos
}

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsGear(r rune) bool {
	return r == '*'
}

func GetGearPos(row int, startPos int, endPos int, sc [][]rune) gearPos {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i == len(sc) {
			continue
		}

		for j := startPos - 1; j <= endPos+1; j++ {
			if j < 0 || j == len(sc[i]) {
				continue
			}

			if IsGear(sc[i][j]) {
				return gearPos{i, j}
			}
		}

	}

	return gearPos{-1, -1}
}

func GetSumGearRatios(s *bufio.Scanner) (int, error) {
	var sc [][]rune
	for s.Scan() {
		sc = append(sc, []rune(s.Text()+"."))
	}

	var gearNumbers []gearNumber
	var gearPositions []gearPos
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

			p := GetGearPos(i, startPos, j-1, sc)
			if p.i != -1 {
				gear := gearNumber{}
				n, err := strconv.Atoi(number)
				if err != nil {
					return 0, err
				}

				gear.number = n
				gear.pos = p
				gearNumbers = append(gearNumbers, gear)

				if !slices.Contains(gearPositions, p) {
					gearPositions = append(gearPositions, p)
				}
			}

			number = ""
			startPos = 0
		}
	}

	totalGearRatios := 0
	for _, p := range gearPositions {
		var partNumbers []int
		for _, gear := range gearNumbers {
			if p == gear.pos {
				partNumbers = append(partNumbers, gear.number)
			}
		}

		if len(partNumbers) <= 1 {
			continue
		}

		gearRatio := 1
		for _, g := range partNumbers {
			gearRatio *= g
		}

		totalGearRatios += gearRatio
	}

	return totalGearRatios, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetSumGearRatios(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
