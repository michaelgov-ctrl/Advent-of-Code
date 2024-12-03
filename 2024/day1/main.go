package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	col1, col2, err := splitFileColumnsToArrays()
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	slices.Sort(col1)
	slices.Sort(col2)

	var totalDistance int
	for i := 0; i < len(col1) && i < len(col2); i++ {
		totalDistance += abs(col1[i] - col2[i])
	}

	fmt.Println("puzzle 1 answer", totalDistance)

	m := make(map[int]int)
	for _, x := range col1 {
		for _, y := range col2 {
			if x == y {
				m[x]++
			}
		}
	}

	var similarityScore int
	for x, y := range m {
		similarityScore += x * y
	}

	fmt.Println("puzzle 2 answer", similarityScore)
}

func splitFileColumnsToArrays() ([]int, []int, error) {
	f, err := os.Open("day1_puzzle_input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("%v", err)
	}
	defer f.Close()

	var col1, col2 []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) < 2 {
			return nil, nil, fmt.Errorf("input string did not match expected format of two space delimited integers: %v", words)
		}

		i1, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse first field to int: %s", words[0])
		}
		col1 = append(col1, i1)

		i2, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse first field to int: %s", words[1])
		}
		col2 = append(col2, i2)
	}

	return col1, col2, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
