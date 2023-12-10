package utils

import "regexp"

func ExtractDigitsFromString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^0-9 ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
