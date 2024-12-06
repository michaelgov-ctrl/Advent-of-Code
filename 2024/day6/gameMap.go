package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrOffMap = errors.New("off map")

type FloorPlan [][]rune

func loadMap() (FloorPlan, Position) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var floorplan FloorPlan
	var pos Position
	var rows int

	for scanner.Scan() {
		text := scanner.Text()
		temp := make([]rune, len(text))
		for i, r := range text {
			if r == '^' {
				pos.X = i
				pos.Y = rows
			}

			temp[i] = r
		}
		floorplan = append(floorplan, temp)

		rows++
	}

	return floorplan, pos
}

func (fp FloorPlan) validate(pos Position) error {
	if pos.Y < 0 || len(fp) <= pos.Y {
		return ErrOffMap
	}

	if pos.X < 0 || len(fp) <= pos.X {
		return ErrOffMap
	}

	return nil
}
