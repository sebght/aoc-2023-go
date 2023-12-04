package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func clearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^0-9 ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func FirstPartDay1(lines []string) int {
	count := 0
	for _, line := range lines {
		digits := clearString(line)
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

func SecondPartDay1(lines []string) int {
	secondCount := 0
	thirdCount := 0

	replaceMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		for word, replacement := range replaceMap {
			line = strings.Replace(line, word, replacement, -1)
		}
		digits := clearString(line)
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

		secondCount += lineValue
		thirdCount++
		fmt.Println("Valeur de la ligne ", thirdCount, " :", lineValue)
	}
	return secondCount
}
