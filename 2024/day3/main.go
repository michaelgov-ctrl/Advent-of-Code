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

type ProcessorState int

const (
	Do ProcessorState = iota
	Dont
)

type Processor struct {
	State ProcessorState
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	// problem 1
	mulRE := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	// problem 2
	mulDoDontRE := regexp.MustCompile(`(mul|do|don't)\((|\d{1,3},\d{1,3})\)`)

	var problemOneSum, problemTwoSum int
	proc := Processor{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		// problem 1
		mulMatches := mulRE.FindAllString(t, -1)
		for _, m := range mulMatches {
			problemOneSum += mustEvaluateMul(m)
		}

		// problem 2
		mulDoDontmatches := mulDoDontRE.FindAllString(t, -1)
		for _, m := range mulDoDontmatches {
			if m == "do()" {
				proc.State = Do
			} else if m == "don't()" {
				proc.State = Dont
			} else {
				if proc.State == Do {
					problemTwoSum += mustEvaluateMul(m)
				}
			}
		}
	}

	fmt.Printf("answer to problem 1: %d\n", problemOneSum)
	fmt.Printf("answer to problem 2: %d\n", problemTwoSum)
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
