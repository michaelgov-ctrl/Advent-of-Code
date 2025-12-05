package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

const REFRESH_RATE = 1 * time.Millisecond

type animator struct {
	area cursor.Area
}

func newAnimator() *animator {
	return &animator{area: cursor.NewArea()}
}

func (a *animator) problem1(path string) int {
	floorplan := newFloorPlan(fromFile(path))
	return a.animateFloorPlanProcessing(floorplan, a.problem1ProcFunc)
}

func (a *animator) problem2(path string) int {
	floorplan := newFloorPlan(fromFile(path))
	return a.animateFloorPlanProcessing(floorplan, a.problem2ProcFunc)
}

func (a *animator) animateFloorPlanProcessing(fp floorplan, pf processFunc) (processedRolls int) {
	cursor.Hide()
	defer cursor.Show()

	a.area.Update(fp.String())

	processedRolls = fp.processPaperRolls(pf)

	a.area.Clear()
	return
}

func (a *animator) problem1ProcFunc(fp floorplan) int {
	var picked int

	validator := func(it item) bool {
		// for problem1 a removed paperroll still counts towards the SURROUNDING_LIMIT
		return it != empty
	}

	a.area.Top()
	for i, row := range fp {
		for j, it := range row {
			if it == paperroll && fp.searchAroundIndex(validator, i, j) < SURROUNDING_LIMIT {
				fp[i][j] = removed
				picked++
				a.updateFloorPlanRow(fp, i)
			}
		}

		a.area.Down(1)
	}

	return picked
}

func (a *animator) problem2ProcFunc(fp floorplan) int {
	picked, found := 0, true

	validator := func(it item) bool {
		// for problem2 we only care about adjacent paperrolls
		return it == paperroll
	}

	// naive search, this got a correct answer, but I wonder if this could miss an optimal case
	for found {
		found = false
		a.area.Top()

		for i, row := range fp {
			for j, it := range row {
				if it == paperroll && fp.searchAroundIndex(validator, i, j) < SURROUNDING_LIMIT {
					fp[i][j] = removed
					picked++
					found = true
					a.updateFloorPlanRow(fp, i)
				}
			}

			a.area.Down(1)
		}
	}

	return picked
}

func (a *animator) updateFloorPlanRow(fp floorplan, i int) {
	a.area.StartOfLine()
	fmt.Print(string(fp[i]))
	time.Sleep(REFRESH_RATE)
}
