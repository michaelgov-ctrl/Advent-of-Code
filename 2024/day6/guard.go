package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
)

type Guard struct {
	UniqueStepCount int
	Direction       Direction
	Position
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X, Y int
}

func (g *game) moveGuardTillOffMapAndRender() {
	for {
		if err := g.moveGuard(); err != nil {
			break
		}

		cursor.Bottom()
		cursor.Up(len(g.floorplan) - g.guard.Y)
		cursor.StartOfLine()
		fmt.Print(string(g.floorplan[g.guard.Y]))

		//g.renderAboveAndBelow()

		time.Sleep(100 * time.Millisecond)
	}
}

func (g *game) renderLine(line int) {
	cursor.StartOfLine()
	fmt.Print(string(g.floorplan[line]))
}

func (g *game) renderAboveAndBelow() {
	cursor.Bottom()
	cursor.Up((len(g.floorplan) - g.guard.Y) - 1)
	for i := 0; i < 3; i++ {
		g.renderLine((len(g.floorplan) - g.guard.Y) + i)
	}
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

func (g *Guard) updatePositionOnFloorPlan(fp FloorPlan) {
	fp[g.Y][g.X] = g.directionSymbol()
}

func (g *Guard) directionSymbol() rune {
	var res = '何'
	switch g.Direction {
	case Up:
		res = '^'
	case Down:
		res = 'v'
	case Left:
		res = '<'
	case Right:
		res = '>'
	}

	return res
}

func (g *Guard) nextPosition() Position {
	pos := g.Position
	switch g.Direction {
	case Up:
		pos.Y--
	case Down:
		pos.Y++
	case Left:
		pos.X--
	case Right:
		pos.X++
	}

	return pos
}

func (pos Position) getTerrain(fp FloorPlan) (rune, error) {
	if err := fp.validate(pos); err != nil {
		return rune(0), err
	}

	return fp[pos.Y][pos.X], nil
}

func (g *Guard) leaveX(fp FloorPlan) {
	fp[g.Y][g.X] = 'X'
}

func (g *Guard) TurnRight() {
	switch g.Direction {
	case Up:
		g.Direction = Right
	case Down:
		g.Direction = Left
	case Left:
		g.Direction = Up
	case Right:
		g.Direction = Down
	}
}
