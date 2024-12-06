package day06

import (
	"log"
)

var guards = map[rune]bool{
	'^': true,
	'>': true,
	'<': true,
	'v': true,
}

func Solve(input []string) {
	m, guardX, guardY := parseInput(input)
	log.Printf("guard x %d y %d", guardX, guardY)
	log.Printf("day 6 part 1: %d\n", part1(m, guardX, guardY))
	log.Printf("day 6 part 2: %d\n", part2(m, guardX, guardY))
}

func part1(m [][]rune, guardX, guardY int) int {
	result := 1
	maxX, maxY := len(m[0]), len(m)

	guard := m[guardY][guardX]

	for {
		x, y := move(guard, guardX, guardY)

		if x == maxX || y == maxY {
			break
		}
		if m[y][x] == '#' {
			guard = rotate(guard)
		} else {
			if m[guardY][guardX] != 'x' {
				result++
			}
			m[guardY][guardX] = 'x'
			guardX, guardY = x, y
		}
	}

	return result
}

func part2(m [][]rune, guardX, guardY int) int {
	return 0
}

func rotate(guard rune) rune {
	switch guard {
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '^':
		return '>'
	default:
		return '^'
	}
}

func move(guard rune, x, y int) (int, int) {
	switch guard {
	case '>':
		return x + 1, y
	case '<':
		return x - 1, y
	case '^':
		return x, y - 1
	// case 'v':
	default:
		return x, y + 1
	}
}

func parseInput(input []string) ([][]rune, int, int) {
	m := make([][]rune, len(input))
	for i := range m {
		m[i] = make([]rune, len(input[i]))
	}

	var guardX, guardY int

	for y, line := range input {
		for x, r := range line {
			m[y][x] = r
			if guards[r] {
				guardX, guardY = x, y
			}
		}
	}

	return m, guardX, guardY
}
