package main

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
