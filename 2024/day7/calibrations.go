package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrBadInput = errors.New("input string not in expected format")

type Operators []rune

type CalibrationTest struct {
	Result    int
	Operands  []int
	Operators Operators
	IsValid   bool
	Solution  string
}

func parseCalibrationTest(str string) (CalibrationTest, error) {
	var ct CalibrationTest

	fields := strings.Fields(str)

	result, err := strconv.Atoi(strings.TrimRight(fields[0], ":"))
	if err != nil {
		return ct, ErrBadInput
	}

	operands := make([]int, len(fields)-1)
	for i, v := range fields[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			return ct, ErrBadInput
		}

		operands[i] = num
	}

	ct.Result = result
	ct.Operands = operands
	ct.Operators = []rune{'+', 'x'}

	return ct, nil
}

func (ct *CalibrationTest) Validate() {
	switch len(ct.Operands) {
	case 0:
		ct.IsValid = false
	case 1:
		b := ct.Operands[0] == ct.Result
		if b {
			ct.Solution = strconv.Itoa(ct.Operands[0])
		}

		ct.IsValid = b
	default:
		ct.IsValid = ct.calculatedResult()
	}
}

func (ct *CalibrationTest) calculatedResult() bool {
	var operations []string
	ct.Operators.generateOperations(len(ct.Operands)-1, "", &operations)

	for _, op := range operations {
		var res = ct.Operands[0]
		for i, v := range op {
			switch v {
			case '+':
				res += ct.Operands[i+1]
			case 'x':
				res *= ct.Operands[i+1]
			default:
				panic(fmt.Sprintf("unexpected operator: %v", v))
			}
		}

		if res == ct.Result {
			ct.Solution = op
			return true
		}
	}

	return false
}

func (o Operators) generateOperations(length int, current string, results *[]string) {
	if len(current) == length {
		*results = append(*results, current)
		return
	}

	for _, char := range o {
		o.generateOperations(length, current+string(char), results)
	}
}
