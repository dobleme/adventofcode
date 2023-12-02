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

func IsPossibleGame(g game, redCubes int, greenCubes int, blueCubes int) bool {
	return g.redCubes <= redCubes &&
		g.greenCubes <= greenCubes &&
		g.blueCubes <= blueCubes
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

func GetPossibleGamesBy(redCubes int, greenCubes int, blueCubes int, s *bufio.Scanner) (int, error) {
	games := []game{}
	for s.Scan() {
		g, err := ParseGame(s.Text())
		if err != nil {
			return 0, err
		}

		if IsPossibleGame(g, redCubes, greenCubes, blueCubes) {
			games = append(games, g)
		}
	}

	totalPossibleGames := 0
	for _, game := range games {
		totalPossibleGames += game.id
	}

	return totalPossibleGames, nil
}

func main() {
	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := GetPossibleGamesBy(redCubes, greenCubes, blueCubes, s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
