package days

import (
	"aoc-2023-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func FirstPartDay1(lines []string) int {
	count := 0
	for _, line := range lines {
		digits := utils.ExtractDigitsFromString(line)
		result := ""
		switch {
		case len(digits) < 2:
			result = digits + digits
		case len(digits) > 2:
			result = string(digits[0]) + string(digits[len(digits)-1])
		default:
			result = digits
		}
		lineValue, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier :", err)
			return 0
		}

		count += lineValue
	}
	return count
}

func find(s string) (int, bool) {
	tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	for k, v := range tokens {
		if strings.Contains(s, k) {
			return v, true
		}
	}
	return 0, false
}

func getFirstDigit(line string) (int, bool) {
	r := 0
	for r < len(line) {
		r++
		value, ok := find(line[:r])
		if ok {
			return value, true
		}
	}
	return 0, false
}

func getLastDigit(line string) (int, bool) {
	l := len(line)
	for l > 0 {
		l--
		value, ok := find(line[l:])
		if ok {
			return value, true
		}
	}
	return 0, false
}

func SecondPartDay1(lines []string) int {
	total := 0

	for _, line := range lines {
		first, _ := getFirstDigit(line)
		last, _ := getLastDigit(line)

		lineValue := first*10 + last
		//fmt.Println("Valeur de la ligne ", i+1, " :", lineValue)
		total += lineValue
	}
	return total
}
