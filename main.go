package main

import (
	"aoc-2023-go/days"
	"aoc-2023-go/utils"
	"fmt"
)

func day1() {
	fmt.Println("=^..^=    =^..^=   Day 1   =^..^=    =^..^=")
	lignes := utils.ReadStringList("inputs/day1.txt")
	fmt.Println("Calibration value :", days.FirstPartDay1(lignes))
	fmt.Println("Second calibration value :", days.SecondPartDay1(lignes))
	fmt.Print("\n")
}

func day2() {
	fmt.Println("=^..^=    =^..^=   Day 2   =^..^=    =^..^=")
	lignes := utils.ReadStringList("inputs/day2.txt")
	fmt.Println("Game Validity Prediction :", days.FirstPartDay2(lignes))
	fmt.Println("Game Total Power :", days.SecondPartDay2(lignes))
	fmt.Print("\n")
}

func main() {
	day2()
}
