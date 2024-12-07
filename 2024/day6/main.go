package main

import (
	"fmt"

	"atomicgo.dev/cursor"
)

type game struct {
	area            cursor.Area
	renderedContent string
	floorplan       FloorPlan
	guard           Guard
}

func main() {
	fp, pos := loadMap()
	var game = game{
		renderedContent: "Advent of Code Day 6\n",
		area:            cursor.NewArea(),
		floorplan:       fp,
		guard: Guard{
			Direction: Up,
			Position:  pos,
		},
	}

	game.area.Update(game.renderedContent)
	for _, row := range game.floorplan {
		game.renderedContent += fmt.Sprintf("%s\n", string(row))
	}
	game.area.Update(game.renderedContent)

	game.moveGuardTillOffMapAndRender()

	var sum = 1
	for _, row := range game.floorplan {
		for _, r := range row {
			if r == 'X' {
				sum++
			}
		}
	}

	fmt.Println("ackshully: ", sum)
	fmt.Println("problem 1 ans: ", game.guard.UniqueStepCount)

}
