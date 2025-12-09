package main

import "fmt"

func main() {
	fmt.Printf("problem 1: %d\n", sumSolutions(parseProblems("day6/_input.txt")))
	fmt.Printf("problem 2: %d\n", sumSolutions(parseCephaloProblems("day6/_input.txt")))
}
