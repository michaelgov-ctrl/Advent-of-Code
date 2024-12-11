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

	var sum int
	for _, ct := range calibrationTests {
		ct.Validate()
		if ct.IsValid {
			sum += ct.Result
		}
	}

	fmt.Println("problem 1 answer: ", sum)
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
