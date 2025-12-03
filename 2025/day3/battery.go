package main

import (
	"strings"
)

type batterybank string

func parseBatteryBank(s string) batterybank {
	return batterybank(s)
}

func (bb batterybank) MaxJoltage(digits int) int {
	var joltage, start int
	for i, step := digits, stepBy(1, digits-1); i > 0; i, step = i-1, step/10 {
		// we always need to look between the index of the last largest digit
		// & an index that buffers to end to fill the full digits worth of joltage
		end := len(bb) - i + 1

		sub := bb[start:end]
		v, subIdx := max(sub)

		start = start + subIdx + 1

		joltage += v * step
	}

	return joltage
}

func max(bb batterybank) (value, index int) {
	max := '\x00'
	for _, r := range bb {
		if r > max {
			max = r
		}
	}

	value, index = int(max-'0'), strings.IndexRune(string(bb), max)
	return
}

func stepBy(n, digits int) int {
	for range digits {
		n *= 10
	}

	return n
}
