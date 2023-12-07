package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	HighCard  = 0
	Pair      = 1
	TwoPair   = 2
	ThreeKind = 3
	FullHouse = 4
	FourKind  = 5
	FiveKind  = 6
)

var valueCards map[rune]int = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type hand struct {
	repr string
	kind int
	bid  int
}

type MinHeap []hand

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].kind == h[j].kind {
		for a, ri := range h[i].repr {
			rj := rune(h[j].repr[a])
			if ri == rj {
				continue
			}

			return valueCards[ri] < valueCards[rj]
		}
	}

	return h[i].kind < h[j].kind
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(hand))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func parseHand(s string) (hand, error) {
	h := hand{repr: s}

	blank := strings.Index(s, " ")
	bid, err := strconv.Atoi(s[blank+1:])
	if err != nil {
		return h, err
	}
	h.bid = bid

	var checkedCards []rune
	kind := HighCard
	jokers := 0

	handKind := s[:blank]
	for _, r := range handKind {
		if r == 'J' {
			jokers++
			continue
		}

		if slices.Contains(checkedCards, r) {
			continue
		}

		c := strings.Count(handKind, string(r))
		switch c {
		case 2:
			kind += Pair
		case 3:
			kind += ThreeKind
		case 4:
			kind += FourKind
		case 5:
			kind += FiveKind
		}

		checkedCards = append(checkedCards, r)
	}

	if jokers > 0 {
		if kind == HighCard {
			kind += (jokers * 2) - 1
		} else {
			kind += jokers * 2
		}
	}

	if kind > FiveKind {
		kind = FiveKind
	}

	h.kind = kind

	return h, nil
}

func GetTotalWinnigns(s []string) (int, error) {
	hands := &MinHeap{}

	for _, h := range s {
		parsedHand, err := parseHand(h)
		if err != nil {
			return 0, err
		}

		heap.Push(hands, parsedHand)
	}

	totalWinnings := 0
	rank := 1
	for hands.Len() > 0 {
		h := heap.Pop(hands).(hand)
		totalWinnings += rank * h.bid
		rank++
	}

	return totalWinnings, nil
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

	res, err := GetTotalWinnigns(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", res)
}
