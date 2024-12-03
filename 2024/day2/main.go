package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	var safeCount, problemDampenedSafeCount int

	pd := ProblemDampener{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums := mustParseLine(scanner.Text())

		pd.State = Off
		if pd.lineIsSafe(nums) {
			safeCount++
		}

		pd.State = On
		if pd.lineIsSafe(nums) {
			problemDampenedSafeCount++
		}
	}

	fmt.Println("puzzle 1 answer is:", safeCount)
	fmt.Println("puzzle 2 answer is:", problemDampenedSafeCount)
}

func mustParseLine(str string) []int {
	fields := strings.Fields(str)

	nums := make([]int, len(fields))
	for i, v := range fields {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("failed to parse string to int: %s", v))
		}

		nums[i] = n
	}

	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
