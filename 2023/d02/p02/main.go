package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id         int
	redCubes   int
	greenCubes int
	blueCubes  int
}

func (g game) Power() int {
	return g.redCubes * g.greenCubes * g.blueCubes
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func ParseGame(s string) (game, error) {
	g := game{}

	parts := strings.Split(s, ":")

	idPart := []rune(parts[0])
	id, err := strconv.Atoi(string(idPart[5:]))
	if err != nil {
		return g, err
	}
	g.id = id

	moves := strings.Split(parts[1], ";")
	for _, move := range moves {
		move = strings.TrimSpace(move)
		cubes := strings.Split(move, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cubeParts := strings.Split(cube, " ")
			totalCubes, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				return g, nil
			}

			switch cubeParts[1] {
			case "red":
				g.redCubes = Max(totalCubes, g.redCubes)
			case "green":
				g.greenCubes = Max(totalCubes, g.greenCubes)
			case "blue":
				g.blueCubes = Max(totalCubes, g.blueCubes)
			default:
				return g, fmt.Errorf("Color cube %s is not valid", cubeParts[1])
			}
		}
	}

	return g, nil
}

func GetSumOfPOwerGames(s *bufio.Scanner) (int, error) {
	games := []game{}
	for s.Scan() {
		g, err := ParseGame(s.Text())
		if err != nil {
			return 0, err
		}

		games = append(games, g)
	}

	totalPowerGames := 0
	for _, game := range games {
		totalPowerGames += game.Power()
	}

	return totalPowerGames, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetSumOfPOwerGames(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
