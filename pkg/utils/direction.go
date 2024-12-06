package utils

const (
	TOP       = 0
	TOP_RIGHT = 1
	RIGHT     = 2
	BOT_RIGHT = 3
	BOT       = 4
	BOT_LEFT  = 5
	LEFT      = 6
	TOP_LEFT  = 7
)

var AllDirections = []Direction{
	TOP,
	TOP_RIGHT,
	RIGHT,
	BOT_RIGHT,
	BOT,
	BOT_LEFT,
	LEFT,
	TOP_LEFT,
}

type Direction int

func (d Direction) Shift(x, y int) (int, int) {
	switch d {
	case TOP:
		y--
	case TOP_RIGHT:
		y--
		x++
	case RIGHT:
		x++
	case BOT_RIGHT:
		y++
		x++
	case BOT:
		y++
	case BOT_LEFT:
		y++
		x--
	case LEFT:
		x--
	case TOP_LEFT:
		y--
		x--
	}
	return x, y
}
