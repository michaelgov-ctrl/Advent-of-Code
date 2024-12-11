package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	calibrationTests, err := loadTests()
	if err != nil {
		log.Fatal(err)
	}

	var problem1Sum, problem2Sum int
	for _, ct := range calibrationTests {
		ct.Validate()
		if ct.IsValid {
			problem1Sum += ct.Result
			problem2Sum += ct.Result
			continue
		}

		ct.Operators = append(ct.Operators, '|')
		ct.Validate()
		if ct.IsValid {
			problem2Sum += ct.Result
		}
	}

	fmt.Println("problem 1 answer: ", problem1Sum)
	fmt.Println("problem 2 answer: ", problem2Sum)
}

func loadTests() ([]CalibrationTest, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var calibrations []CalibrationTest

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ct, err := parseCalibrationTest(scanner.Text())
		if err != nil {
			return nil, err
		}

		calibrations = append(calibrations, ct)
	}

	return calibrations, nil
}
