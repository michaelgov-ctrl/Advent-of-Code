package main

import (
	"fmt"

	"github.com/michaelgov-ctrl/aoc/2025/utils"
)

func main() {
	tm := NewTachyonManifold("day7/_input.txt")
	tm.ProcessTachyon()
	fmt.Println(tm)
}

type status rune

const (
	start    status = 'S'
	splitter status = '^'
	empty    status = '.'
	beam     status = '|'
)

type TachyonManifold struct {
	overview   [][]status
	beamSplits int
}

func NewTachyonManifold(path string) TachyonManifold {
	var tm TachyonManifold
	utils.ForEachLineInFile(path, func(s string) {
		tm.overview = append(tm.overview, []status(s))
	})

	return tm
}

func (nt TachyonManifold) String() string {
	var s string
	for _, row := range nt.overview {
		s += fmt.Sprintf("%s\n", string(row))
	}

	s += fmt.Sprintf("# of beam splits: %d\n", nt.beamSplits)

	return s
}

func (nt *TachyonManifold) ProcessTachyon() {
	nt.beamSplits = 0

	height := len(nt.overview)
	width := len(nt.overview[0]) // all rows are same length

	for i := range height {
		for j := range width {
			switch nt.overview[i][j] {
			case start:
				// always on first line
				nt.overview[i+1][j] = beam

			case splitter:
				if !nt.BeamAbove(i, j) {
					continue
				}

				if j-1 >= 0 {
					nt.overview[i][j-1] = beam
				}

				if j+1 < width {
					nt.overview[i][j+1] = beam
				}

				nt.beamSplits++

			case empty:
				if !nt.BeamAbove(i, j) {
					continue
				}

				nt.overview[i][j] = beam
			}
		}
	}
}

func (nt TachyonManifold) BeamAbove(i, j int) bool {
	if i-1 < 0 {
		return false
	}

	return nt.overview[i-1][j] == beam
}
