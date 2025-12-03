package main

import (
	"fmt"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

func main() {
	fmt.Printf("problem 1: %d\n", problem1("day3/_input.txt"))
	fmt.Printf("problem 2: %d\n", problem2("day3/_input.txt"))
}

func problem1(path string) int {
	return commonLogic(path, 2)
}

func problem2(path string) int {
	return commonLogic(path, 12)
}

func commonLogic(path string, digits int) int {
	var sum int
	utils.ForEachLineInFile(path, func(s string) {
		sum += parseBatteryBank(s).MaxJoltage(digits)
	})

	return sum
}
