package day04

import "log"

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

var allDirections = []Direction{
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

func (d Direction) isUp() bool {
	return d == TOP || d == TOP_RIGHT || d == TOP_LEFT
}

func (d Direction) isDown() bool {
	return d == BOT_LEFT || d == BOT || d == BOT_RIGHT
}

func (d Direction) isRight() bool {
	return d == TOP_RIGHT || d == RIGHT || d == BOT_RIGHT
}

func (d Direction) isLeft() bool {
	return d == TOP_LEFT || d == LEFT || d == BOT_LEFT
}

func (d Direction) shift(x, y int) (int, int) {
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

type Field [][]rune

func NewField(input []string) Field {
	field := make([][]rune, len(input))
	for i, line := range input {
		field[i] = []rune(line)
	}
	return field
}

func (f *Field) maxY() int {
	return len(*f) - 1
}

func (f *Field) maxX() int {
	return len((*f)[0]) - 1
}

func (f *Field) has(what rune, x, y int, direction Direction) bool {
	if direction.isUp() && y == 0 {
		return false
	}

	if direction.isDown() && y == f.maxY() {
		return false
	}

	if direction.isRight() && x == f.maxX() {
		return false
	}

	if direction.isLeft() && x == 0 {
		return false
	}

	x, y = direction.shift(x, y)

	return (*f)[y][x] == what
}
func (f *Field) hasXmas(x, y int, direction Direction) bool {
	if (*f)[y][x] != 'X' {
		return false
	}
	if !f.has('M', x, y, direction) {
		return false
	}

	x, y = direction.shift(x, y)
	if !f.has('A', x, y, direction) {
		return false
	}

	x, y = direction.shift(x, y)
	if !f.has('S', x, y, direction) {
		return false
	}

	return true
}

func (f *Field) hasXdMases(x, y int, direction Direction) bool {
	if (*f)[y][x] != 'X' {
		return false
	}
	if !f.has('M', x, y, direction) {
		return false
	}

	x, y = direction.shift(x, y)
	if !f.has('A', x, y, direction) {
		return false
	}

	x, y = direction.shift(x, y)
	if !f.has('S', x, y, direction) {
		return false
	}

	return true
}

func Solve(input []string) {
	field := NewField(input)

	log.Printf("day 4 part 1: %d\n", part1(&field))
	log.Printf("day 4 part 2: %d\n", part2(&field))
}

func part1(field *Field) int {
	var result int

	for y, line := range *field {
		for x := range line {
			for _, d := range allDirections {
				if field.hasXmas(x, y, d) {
					result++
				}
			}
		}
	}
	return result
}
func part2(field *Field) int {
	var result int

	for y, line := range *field {
		for x := range line {
			for _, d := range allDirections {
				if field.hasXdMases(x, y, d) {
					result++
				}
			}
		}
	}
	return result
}
