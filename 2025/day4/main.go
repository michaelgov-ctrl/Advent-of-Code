package main

import (
	"fmt"
)

func main() {
	animator := newAnimator()
	fmt.Printf("problem 1: %d\n", animator.problem1("day4/_input.txt"))
	fmt.Printf("problem 2: %d\n", animator.problem2("day4/_input.txt"))
}
