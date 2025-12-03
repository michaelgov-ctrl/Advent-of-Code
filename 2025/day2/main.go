package main

import (
	"fmt"
	"strings"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

func main() {
	fmt.Printf("problem 1: %d\n", problem1("day2/_input.txt"))
	fmt.Printf("problem 2: %d\n", problem2("day2/_input.txt"))
}

func problem1(path string) int {
	return commonLogic(path, isValidId)
}

func problem2(path string) int {
	return commonLogic(path, isValidStrictId)
}

func commonLogic(path string, valid func(int) bool) int {
	var sum int
	utils.ForEachLineInFile(path, func(s string) {
		for r := range strings.SplitSeq(s, ",") {
			start, end, err := parseIdRange(r)
			if err != nil {
				panic(fmt.Sprintf("failed on range: %s, with err: %v", r, err))
			}

			for n := range idRangeSeq(start, end) {
				if !valid(n) {
					sum += n
				}
			}
		}
	})

	return sum
}
