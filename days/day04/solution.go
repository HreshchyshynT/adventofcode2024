package day04

import (
	"log"
)

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

func (f *Field) contains(x, y int) bool {
	return x >= 0 && x <= f.maxX() && y >= 0 && y <= f.maxY()
}

func (f *Field) hasWord(x, y int, word []rune, direction Direction) bool {
	for _, r := range word {
		if !f.contains(x, y) || (*f)[y][x] != r {
			return false
		}
		x, y = direction.shift(x, y)
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
	word := []rune{'X', 'M', 'A', 'S'}

	for y, line := range *field {
		for x := range line {
			for _, d := range allDirections {
				if field.hasWord(x, y, word, d) {
					result++
				}
			}
		}
	}
	return result
}
func part2(field *Field) int {
	var result int

	mas, sam := []rune{'M', 'A', 'S'}, []rune{'S', 'A', 'M'}

	for y := 0; y <= field.maxY()-2; y++ {
		for x := 0; x <= field.maxX()-2; x++ {

			hasMas := field.hasWord(x, y, mas, BOT_RIGHT)
			hasSam := field.hasWord(x, y, sam, BOT_RIGHT)
			if !hasMas && !hasSam {
				continue
			}
			hasSam = field.hasWord(x+2, y, sam, BOT_LEFT)
			hasMas = field.hasWord(x+2, y, mas, BOT_LEFT)
			if hasMas || hasSam {
				result++
			}
		}
	}
	return result
}
