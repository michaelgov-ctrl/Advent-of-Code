package main

import "strconv"

/*
	Just because I can abstract doesn't mean I should...
*/

type direction int

const (
	left direction = iota
	right
	unknown
)

func mustParseDirection(s string) direction {
	switch s {
	case "L":
		return left
	case "R":
		return right
	default:
		panic("invalid direction")
	}
}

func (d direction) String() string {
	switch d {
	case left:
		return "Left"
	case right:
		return "Right"
	default:
		return "Unknown"
	}
}

type dial struct {
	current     int
	max         int
	min         int
	wrapArounds int
}

func newDial(curr, max, min int) dial {
	return dial{
		current: curr,
		max:     max,
		min:     min,
	}
}

func (d *dial) turn(input string) {
	direction, distance := parseTurn(input)

	switch direction {
	case left:
		d.dec(distance)
	case right:
		d.inc(distance)
	}
}

func (d *dial) inc(n int) {
	/*
		// TODO: remember how to do this with math not iteration??
		remainder := n - (d.max - d.current)
		if remainder > 0 {

		}
	*/

	temp := d.current
	for range n {
		temp++

		if temp > d.max {
			temp = d.min
		}

		if temp == d.min {
			d.wrapArounds++
		}
	}
	d.current = temp
}

func (d *dial) dec(n int) {
	// TODO: math not iteration
	temp := d.current
	for range n {
		temp--

		if temp == d.min {
			d.wrapArounds++
		}

		if temp < d.min {
			temp = d.max
		}
	}
	d.current = temp
}

func (d dial) atMin() bool {
	return d.min == d.current
}

func parseTurn(input string) (direction, int) {
	dir := mustParseDirection(input[:1])

	distance, err := strconv.Atoi(input[1:])
	if err != nil {
		panic("invalid turn")
	}

	return dir, distance
}
