package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var words []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
var wordNumber map[string]rune = map[string]rune{
	words[0]: '1',
	words[1]: '2',
	words[2]: '3',
	words[3]: '4',
	words[4]: '5',
	words[5]: '6',
	words[6]: '7',
	words[7]: '8',
	words[8]: '9',
}

func isDigit(r rune) bool {
	return r == '0' || r == '1' || r == '2' ||
		r == '3' || r == '4' || r == '5' ||
		r == '6' || r == '7' || r == '8' || r == '9'
}

func isNumberCharacter(r rune) bool {
	return r == 'o' || r == 'n' || r == 'e' || r == 't' || r == 'w' ||
		r == 'h' || r == 'r' || r == 'f' || r == 'u' || r == 'i' ||
		r == 'v' || r == 's' || r == 'x' || r == 'g'
}

func isLastRuneNumberWord(r rune) bool {
	return r == 'e' || r == 'o' || r == 'r' || r == 'n' || r == 't' || r == 'f' || r == 's'
}

func clearPaddingWord(r []rune) []rune {
	for i := 1; i < len(r); i++ {
		pre := r[i:len(r)]

		for _, numberWord := range words {
			if strings.HasPrefix(numberWord, string(pre)) {
				return pre
			}
		}
	}

	return []rune{}
}

func startsWithNumberWord(s string) bool {
	for _, numberWord := range words {
		if strings.HasPrefix(numberWord, s) {
			return true
		}
	}

	return false
}

func SumNumbersByLine(s *bufio.Scanner) (int, error) {
	total := 0
	number := make([]rune, 2)
	var word []rune

	j := 1
	for s.Scan() {
		i := 0
		clear(number)
		word = []rune{}

		for _, r := range s.Text() {
			if isDigit(r) {
				number[i] = r
				if i == 0 {
					i++
				}
				word = []rune{}
				continue
			}

			if !isNumberCharacter(r) {
				word = []rune{}
				continue
			}

			word = append(word, r)

			if !startsWithNumberWord(string(word)) {
				word = clearPaddingWord(word)
				continue
			}

			n, ok := wordNumber[string(word)]
			if ok {
				number[i] = n
				word = clearPaddingWord(word)
				if i == 0 {
					i++
				}
			}

		}

		if number[1] == 0 {
			number[1] = number[0]
		}

		n, err := strconv.Atoi(string(number))
		if err != nil {
			return total, err
		}
		j++

		total += n
	}

	return total, nil
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	res, err := SumNumbersByLine(s)

	fmt.Printf("%d\n", res)
}
