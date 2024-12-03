package main

import "slices"

type ProblemDampenerState int

const (
	Off ProblemDampenerState = iota
	On
)

type ProblemDampener struct {
	State ProblemDampenerState
}

func (p *ProblemDampener) lineIsSafe(nums []int) bool {
OUTER:
	for i := 0; i < len(nums); i++ {
		sharedNums := nums

		if p.State == On {
			sharedNums = slices.Delete(slices.Clone(nums), i, i+1)
		}

		var inc, dec bool
		for j := 1; j < len(sharedNums); j++ {
			diff := abs(sharedNums[j-1] - sharedNums[j])

			if !(1 <= diff && diff <= 3) {
				if p.State == On {
					continue OUTER
				}

				return false
			}

			if sharedNums[j-1] > sharedNums[j] {
				dec = true
			} else if sharedNums[j-1] < sharedNums[j] {
				inc = true
			}

			if inc && dec {
				if p.State == On {
					continue OUTER
				}

				return false
			}
		}

		return true
	}

	return false
}
