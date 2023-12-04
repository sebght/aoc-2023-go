package main

import (
	"aoc-2023-go/days"
	"aoc-2023-go/utils"
	"fmt"
)

func day1() {
	fmt.Println("=^..^=    =^..^=   Day 1   =^..^=    =^..^=")
	lignes := utils.ReadStringList("inputs/day1.txt")
	//fmt.Println("Calibration value :", days.FirstPartDay1(lines))
	fmt.Print("\n")
	fmt.Println("Second calibration value :", days.SecondPartDay1(lignes))
}

func main() {
	day1()
}
