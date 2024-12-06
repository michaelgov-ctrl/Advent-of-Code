package main

import (
	"fmt"
)

func main() {
	precedence := getPrecedence()
	pageSets := getPageSets()

	var problem1Sum, problem2Sum int
	for _, ps := range pageSets {
		if ps.isOrdered(precedence) {
			problem1Sum += ps.MiddleValue()
		} else {
			ps.Order(precedence)
			problem2Sum += ps.MiddleValue()
		}
	}

	fmt.Println("Problem 1 answer: ", problem1Sum)
	fmt.Println("Problem 2 answer: ", problem2Sum)
}
