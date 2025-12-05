package main

import (
	"fmt"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

const SURROUNDING_LIMIT = 4

type floorplan [][]item

func newFloorPlan(opts ...Option) floorplan {
	var fp floorplan

	// this should be ordered since fromFile mutates the whole thing...
	for _, opt := range opts {
		opt(&fp)
	}

	return fp
}

type Option func(*floorplan)

func fromFile(path string) Option {
	return func(fp *floorplan) {
		var i int
		utils.ForEachLineInFile(path, func(s string) {
			l := len(s)
			*fp = append(*fp, make([]item, l))
			for j, r := range s {
				(*fp)[i][j] = runeToItem(r)
			}
			i++
		})
	}
}

func (fp floorplan) processPaperRolls(pf processFunc) (processedRolls int) {
	processedRolls = pf(fp)
	return
}

// this could be named better
func (fp floorplan) searchAroundIndex(sf searchFunc, i, j int) int {
	if (i < 0 || i >= len(fp)) || (j < 0 || j >= len(fp[i])) {
		// index out of bounds
		return 0
	}

	var instances int
	for k := i - 1; k <= i+1; k++ {
		if k < 0 || k >= len(fp) {
			// point out of bounds
			continue
		}

		for l := j - 1; l <= j+1; l++ {
			if l < 0 || l >= len(fp[k]) {
				// point out of bounds
				continue
			}

			if k == i && l == j {
				// skip centeral index
				continue
			}

			if sf(fp[k][l]) {
				instances++
			}
		}
	}

	return instances
}

func (fp floorplan) String() string {
	var s string
	for _, row := range fp {
		s += fmt.Sprintf("%s\n", string(row))
	}
	return s
}

type processFunc func(fp floorplan) int

type searchFunc func(it item) bool

type item rune

const (
	paperroll item = '@'
	removed   item = 'X'
	empty     item = '.'
)

func runeToItem(r rune) item {
	switch r {
	case '@':
		return paperroll
	case 'X':
		return removed
	case '.':
		return empty
	default:
		panic(fmt.Sprintf("invalid rune to item conversion: %c", r))
	}
}
