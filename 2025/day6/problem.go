package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

type mathType int

const (
	columnar mathType = iota
	cephalo_columnar
)

type problem struct {
	operands []int
	operator operator
	mathType mathType
}

/*
// this is a lazy string solution
// i'm sure there is a %10 method to math it out
func (p *problem) toCephaloColumnar() {
	if p.mathType == cephalo_columnar {
		return
	}

	p.mathType = cephalo_columnar

	strs, max := make([]string, len(p.operands)), 0
	for i, n := range p.operands {
		s := strconv.Itoa(n)
		strs[i] = s
		if len(s) > max {
			max = len(s)
		}
	}

	for i := range strs {
		strs[i] = fmt.Sprintf("%*s", max, strs[i]) // pad to max
	}

	cols := make([]string, max)
	for i := max - 1; i >= 0; i-- {
		for _, s := range strs {
			if len(s)-1 < i || s[i] == ' ' { // bounds check shouldnt be necessary..
				continue
			}

			cols[i] += string(s[i])
		}
	}

	operands := make([]int, len(cols))
	for i, n := range cols {
		if n == "" {
			continue
		}

		op, err := strconv.Atoi(n)
		if err != nil {
			panic("how'd we not get a number from a number")
		}

		operands[i] = op
	}

	p.operands = operands
}
*/

func (p problem) solve() int {
	if len(p.operands) == 0 {
		return 0
	}

	var res int
	operation := func(n int) {}

	switch p.operator {
	case add:
		operation = func(n int) {
			res += n
		}
	case multiply:
		res = 1
		operation = func(n int) {
			res *= n
		}
	}

	for _, n := range p.operands {
		operation(n)
	}

	return res
}

func sumSolutions(problems []problem) int {
	var sum int
	for _, p := range problems {
		sum += p.solve()
	}

	return sum
}

func parseProblems(path string) []problem {
	// this is so lazy
	var l int
	utils.ForEachLineInFile(path, func(s string) {
		l = len(strings.Fields(s))
	})

	problems := make([]problem, l)
	utils.ForEachLineInFile(path, func(s string) {
		for i, s := range strings.Fields(s) {
			op := parseOperator([]rune(s)[0])
			if op != unknown {
				problems[i].operator = op
				problems[i].mathType = columnar
				continue
			}

			n, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("%v: %s", err, s))
			}

			problems[i].operands = append(problems[i].operands, n)
		}
	})

	return problems
}

// this is chatgpt, I need to deconstruct and consider
// how to split out the partially padded "chunks"
func parseCephaloProblems(path string) []problem {
	// 1. Read all lines
	var lines []string
	utils.ForEachLineInFile(path, func(s string) {
		lines = append(lines, s)
	})
	if len(lines) == 0 {
		return nil
	}

	h := len(lines)

	// 2. Pad all lines to same width
	w := 0
	for _, s := range lines {
		if len(s) > w {
			w = len(s)
		}
	}

	grid := make([][]rune, h)
	for i, s := range lines {
		row := []rune(s)
		if len(row) < w {
			pad := make([]rune, w-len(row))
			for j := range pad {
				pad[j] = ' '
			}
			row = append(row, pad...)
		}
		grid[i] = row
	}

	// 3. Find problem column ranges: groups of non-empty columns
	type rng struct{ start, end int }
	var ranges []rng
	in := false
	start := 0
	for c := 0; c < w; c++ {
		empty := true
		for r := 0; r < h; r++ {
			if grid[r][c] != ' ' {
				empty = false
				break
			}
		}
		if !in && !empty {
			in = true
			start = c
		} else if in && empty {
			ranges = append(ranges, rng{start, c - 1})
			in = false
		}
	}
	if in {
		ranges = append(ranges, rng{start, w - 1})
	}

	// 4. Build problems: for each range, read operator + column-numbers
	bottom := h - 1
	var probs []problem

	for _, seg := range ranges {
		// operator: scan bottom row in this segment
		op := unknown
		for c := seg.start; c <= seg.end; c++ {
			ch := grid[bottom][c]
			if ch == '+' {
				op = add
				break
			}
			if ch == '*' {
				op = multiply
				break
			}
		}
		if op == unknown {
			continue
		}

		var operands []int

		// Part 2: one number per column, digits top→bottom (ignoring spaces)
		for c := seg.start; c <= seg.end; c++ {
			var sb strings.Builder
			for r := 0; r < bottom; r++ {
				ch := grid[r][c]
				if ch >= '0' && ch <= '9' {
					sb.WriteRune(ch)
				}
			}
			if sb.Len() == 0 {
				continue
			}
			n, err := strconv.Atoi(sb.String())
			if err != nil {
				panic(err)
			}
			operands = append(operands, n)
		}

		probs = append(probs, problem{
			operands: operands,
			operator: op,
			mathType: cephalo_columnar,
		})
	}

	return probs
}

type operator rune

const (
	add      operator = '+'
	multiply operator = '*'
	unknown  operator = '\x00'
)

func parseOperator(r rune) operator {
	switch r {
	case rune(add):
		return add
	case rune(multiply):
		return multiply
	default:
		return unknown
	}
}
