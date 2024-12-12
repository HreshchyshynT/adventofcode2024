package day12

import (
	"adventofcode2024/pkg/utils"
	"log"
)

func Solve(input []string) {

	log.Printf("day 12 part 1: %d", part1(parseInput(input)))
}

var directions = []utils.Direction{
	utils.TOP,
	utils.RIGHT,
	utils.BOT,
	utils.LEFT,
}

func part1(input [][]rune) int {
	result := 0
	for i, line := range input {
		for j, r := range line {
			if r != '.' {
				a, p := findAreaAndPerimeter(input, j, i, make(map[point]bool))
				result += a * p
			}
		}
	}
	return result
}

type point struct {
	x, y int
}

func findAreaAndPerimeter(input [][]rune, x, y int, passed map[point]bool) (a, p int) {
	symbol := input[y][x]
	passed[point{x: x, y: y}] = true
	a += 1
	for _, d := range directions {
		dx, dy := d.Shift(x, y)
		po := point{x: dx, y: dy}
		if passed[po] {
			continue
		}
		candidate := '.'

		if dy >= 0 && dy < len(input) && dx >= 0 && dx < len(input[dy]) {
			candidate = input[dy][dx]
		}
		if candidate != '.' && candidate == symbol {
			aa, pp := findAreaAndPerimeter(input, dx, dy, passed)
			a += aa
			p += pp
		} else {
			p += 1
		}
	}
	input[y][x] = '.'

	return a, p
}

func parseInput(input []string) [][]rune {
	result := make([][]rune, len(input))
	for i, line := range input {
		result[i] = []rune(line)
	}

	return result
}
