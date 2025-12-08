package main

import "fmt"

func main() {
	ranges := reduceRanges(parseIdRanges("day5/_input_id_ranges.txt"))
	ingredients := parseIngredients("day5/_input_ids.txt")

	fmt.Printf("problem 1: %d\n", problem1(ranges, ingredients))
	fmt.Printf("problem 2: %d\n", problem2(ranges))
}

func problem1(ranges []idRange, ingredients []ingredient) int {
	var sum int
	for _, ingr := range ingredients {
		if ingr.isFresh(ranges) {
			sum++
		}
	}
	return sum
}

func problem2(ranges []idRange) int {
	var sum int
	for _, r := range ranges {
		sum += r.end - r.start + 1 // range is inclusive, add 1
	}
	return sum
}
