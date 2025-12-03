package main

import (
	"fmt"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

func main() {
	fmt.Printf("problem 1: %d\n", problem1("day1/_input.txt"))
	fmt.Printf("problem 2: %d\n", problem2("day1/_input.txt"))
}

func problem1(path string) int {
	dial, minCounter := newDial(50, 99, 0), 0
	utils.ForEachLineInFile(path, func(s string) {
		dial.turn(s)
		if dial.atMin() {
			minCounter++
		}
	})

	return minCounter
}

func problem2(path string) int {
	dial := newDial(50, 99, 0)
	utils.ForEachLineInFile(path, func(s string) {
		dial.turn(s)
	})

	return dial.wrapArounds
}
