package days

import (
	"strconv"
	"strings"
)

func isDrawValid(draw string, limits []int) bool {
	elements := strings.Split(draw, ", ")
	for _, el := range elements {
		n, color, _ := strings.Cut(el, " ")
		//fmt.Println("color :", color, " et nombre :", n)
		number, _ := strconv.Atoi(n)
		switch color {
		case "red":
			if number > limits[0] {
				return false
			}
		case "green":
			if number > limits[1] {
				return false
			}
		case "blue":
			if number > limits[2] {
				return false
			}
		default:
			//fmt.Println("case :", color)
		}
	}
	return true
}

func isGameValid(draws []string, limits []int) bool {
	for _, draw := range draws {
		if !isDrawValid(draw, limits) {
			return false
		}
	}
	return true
}

func FirstPartDay2(lines []string) int {
	count := 0
	limits := []int{12, 13, 14}

	for _, line := range lines {
		split := strings.SplitN(line, " ", 3)
		gameNumber, _ := strconv.Atoi(strings.Trim(split[1], ":"))
		draws := strings.Split(split[2], "; ")

		if isGameValid(draws, limits) {
			//fmt.Println("Jeu valide numéro :", gameNumber)
			count += gameNumber
		}
	}
	return count
}

func computeGameTotalPower(line string) int {
	split := strings.SplitN(line, " ", 3)
	draws := strings.Split(split[2], "; ")
	lowests := map[string]int{"red": 0, "blue": 0, "green": 0}
	for _, draw := range draws {
		elements := strings.Split(draw, ", ")
		for _, el := range elements {
			n, color, _ := strings.Cut(el, " ")
			//fmt.Println("color :", color, " et nombre :", n)
			number, _ := strconv.Atoi(n)
			if number > lowests[color] {
				lowests[color] = number
			}
		}
	}
	return lowests["red"] * lowests["blue"] * lowests["green"]
}

func SecondPartDay2(lines []string) int {
	count := 0
	for _, line := range lines {
		count += computeGameTotalPower(line)
	}
	return count
}
