package main

import (
	"fmt"
)

func main() {
	pageSets := getPageSets()
	pagePrecedence := getPagePrecedence()

	var problem1Sum, problem2Sum int
	for _, ps := range pageSets {
		if ps.isOrdered(pagePrecedence) {
			problem1Sum += ps.MiddleValue()
		} else {
			ps.Order(pagePrecedence)
			problem2Sum += ps.MiddleValue()
		}
	}

	fmt.Println("Problem 1 answer: ", problem1Sum)
	fmt.Println("Problem 2 answer: ", problem2Sum)
}
