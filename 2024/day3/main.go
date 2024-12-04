package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		for _, m := range matches {
			sum += mustEvaluateMul(m)
		}
	}

	fmt.Printf("answer to problem 1: %d\n", sum)
}

func mustEvaluateMul(str string) int {
	stripped := strings.TrimSuffix(strings.TrimPrefix(str, "mul("), ")")
	nums := strings.Split(stripped, ",")

	if len(nums) != 2 {
		panic(fmt.Sprintf("mul expression contained more than 2 inputs: %v", nums))
	}

	var product = 1
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(fmt.Sprintf("failed to parse string to int: %v", num))
		}

		product *= n
	}

	return product
}
