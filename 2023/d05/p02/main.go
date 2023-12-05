package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Source      = 1
	Destination = 0
	Length      = 2
)

type categoryMap struct {
	source      int
	destination int
	length      int
}

func convertIntoIntSlice(s string) ([]int, error) {
	sNumbers := strings.Split(s, " ")
	n := make([]int, len(sNumbers))

	for i, sn := range sNumbers {
		number, err := strconv.Atoi(sn)
		if err != nil {
			return nil, err
		}

		n[i] = number
	}

	return n, nil
}

func resolveDestination(c []categoryMap, source int) int {
	if source < c[0].source || source >= c[len(c)-1].source+c[len(c)-1].length {
		return source
	}

	for _, m := range c {
		if source >= m.source && source < m.source+m.length {
			return m.destination + (source - m.source)
		}
	}

	return source
}

// Im assuming the whole file is in perfect order
func GetMinLocationSeed(s *bufio.Scanner) (int, error) {
	s.Scan()
	line := s.Text()
	pairSeeds, err := convertIntoIntSlice(line[7:])
	if err != nil {
		return 0, err
	}

	s.Scan()

	i := 0
	var seedLocationMap [][]categoryMap
	for s.Scan() {
		line = s.Text()
		if line == "" {
			sort.Slice(seedLocationMap[i], func(a, b int) bool {
				return seedLocationMap[i][a].source < seedLocationMap[i][b].source
			})
			i++
			continue
		}

		if strings.Contains(line, "map") {
			seedLocationMap = append(seedLocationMap, []categoryMap{})
			continue
		}

		m, err := convertIntoIntSlice(line)
		if err != nil {
			return 0, err
		}
		seedLocationMap[i] = append(
			seedLocationMap[i],
			categoryMap{
				source:      m[Source],
				destination: m[Destination],
				length:      m[Length],
			},
		)
	}

	minLocation := math.MaxInt
	seed := 0
	for i := 0; i < len(pairSeeds); i += 2 {
		for j := 0; j < pairSeeds[i+1]; j++ {
			seed = j + pairSeeds[i]

			for _, m := range seedLocationMap {
				seed = resolveDestination(m, seed)
			}

			if seed < minLocation {
				minLocation = seed
			}
		}
	}

	return minLocation, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	res, err := GetMinLocationSeed(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
