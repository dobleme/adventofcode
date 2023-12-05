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
	var si []int
	for _, n := range strings.Split(s, " ") {
		seed, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}

		si = append(si, seed)
	}

	return si, nil
}

func resolveDestination(c []categoryMap, source int) int {
	if source < c[0].source {
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
	seeds, err := convertIntoIntSlice(line[7:])
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
			categoryMap{source: m[Source], destination: m[Destination], length: m[Length]},
		)
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		source := seed
		for _, m := range seedLocationMap {
			source = resolveDestination(m, source)
		}

		if source < minLocation {
			minLocation = source
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
