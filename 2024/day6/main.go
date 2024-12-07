package main

import (
	"fmt"
)

type game struct {
	renderedContent string
	floorplan       FloorPlan
	guard           Guard
}

func main() {
	fp, pos := loadMap()
	var game = game{
		renderedContent: "Advent of Code Day 6\n",
		floorplan:       fp,
		guard: Guard{
			UniqueStepCount: 1,
			Direction:       Up,
			Position:        pos,
		},
	}

	for _, row := range game.floorplan {
		game.renderedContent += fmt.Sprintf("%s\n", string(row))
	}
	fmt.Print(game.renderedContent)

	game.moveGuardTillOffMapAndRender()

	fmt.Println("problem 1 answer: ", game.guard.UniqueStepCount)
}
