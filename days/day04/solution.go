package day04

import (
	"adventofcode2024/pkg/utils"
	"log"
)

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

func (f *Field) hasWord(x, y int, word []rune, direction utils.Direction) bool {
	for _, r := range word {
		if !f.contains(x, y) || (*f)[y][x] != r {
			return false
		}
		x, y = direction.Shift(x, y)
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
			for _, d := range utils.AllDirections {
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

			hasMas := field.hasWord(x, y, mas, utils.BOT_RIGHT)
			hasSam := field.hasWord(x, y, sam, utils.BOT_RIGHT)
			if !hasMas && !hasSam {
				continue
			}
			hasSam = field.hasWord(x+2, y, sam, utils.BOT_LEFT)
			hasMas = field.hasWord(x+2, y, mas, utils.BOT_LEFT)
			if hasMas || hasSam {
				result++
			}
		}
	}
	return result
}
