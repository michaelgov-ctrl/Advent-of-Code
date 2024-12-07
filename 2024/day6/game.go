package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

func (g *game) moveGuardTillOffMapAndRender() {
	for {
		if err := g.moveGuard(); err != nil {
			break
		}

		location := len(g.floorplan) - g.guard.Y
		cursor.Bottom()
		g.renderLine(location)

		switch g.guard.Direction {
		case Up:
			cursor.Down(1)
			cursor.StartOfLine()
			fmt.Print(string(g.floorplan[g.guard.Y+1]))
		case Down:
			cursor.Up(1)
			cursor.StartOfLine()
			fmt.Print(string(g.floorplan[g.guard.Y-1]))
		}

		time.Sleep(2 * time.Millisecond)

	}

	cursor.Bottom()
}

func (g *game) moveGuard() error {
	pos := g.guard.nextPosition()
	terrain, err := pos.getTerrain(g.floorplan)
	if err != nil {
		return err
	}

	switch terrain {
	case '#':
		g.guard.TurnRight()
	case '.':
		g.guard.UniqueStepCount++

		g.guard.leaveX(g.floorplan)
		g.guard.Position = pos
	default:
		g.guard.leaveX(g.floorplan)
		g.guard.Position = pos
	}

	g.guard.updatePositionOnFloorPlan(g.floorplan)

	return nil
}

func (g *game) renderLine(line int) {
	cursor.Move(0, line)
	cursor.ClearLine()
	fmt.Print(string(g.floorplan[g.guard.Y]))
}
