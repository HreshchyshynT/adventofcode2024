package day10

import (
	"adventofcode2024/pkg/utils"
	"log"
)

type point struct {
	x, y int
}

var availableDirections = []utils.Direction{
	utils.TOP,
	utils.BOT,
	utils.LEFT,
	utils.RIGHT,
}

func Solve(input []string) {

	log.Printf("day 10 part 1: %d", part1(parseInput(input)))
	log.Printf("day 10 part 2: %d", part2(parseInput(input)))
}

func parseInput(input []string) [][]rune {
	result := make([][]rune, len(input))
	for i, line := range input {
		result[i] = []rune(line)
	}
	return result
}

func part1(input [][]rune) int {
	var result int
	for y, line := range input {
		for x, r := range line {
			height := utils.IntParseRune(r)
			if height > 0 {
				continue
			}
			ends := getHickingTrailsEnds(input, x, y)

			result += uniqueElements(ends)
		}
	}

	return result
}

func part2(input [][]rune) int {
	var result int
	for y, line := range input {
		for x, r := range line {
			height := utils.IntParseRune(r)
			if height > 0 {
				continue
			}
			result += getHickingTrailsCount(input, x, y)
		}
	}

	return result
}

func getHickingTrailsCount(runes [][]rune, x, y int) int {
	var result int
	currentHeight := utils.IntParseRune(runes[y][x])
	if currentHeight == 9 {
		return 1
	}

	for _, direction := range availableDirections {
		xx, yy := direction.Shift(x, y)
		if !inRange(runes, xx, yy) {
			continue
		}
		height := utils.IntParseRune(runes[yy][xx])
		if height-currentHeight == 1 {
			count := getHickingTrailsCount(runes, xx, yy)
			result += count
		}
	}
	return result
}

func getHickingTrailsEnds(runes [][]rune, x, y int) []point {
	result := make([]point, 0)
	currentHeight := utils.IntParseRune(runes[y][x])
	if currentHeight == 9 {
		return []point{{y: y, x: x}}
	}

	for _, direction := range availableDirections {
		xx, yy := direction.Shift(x, y)
		if !inRange(runes, xx, yy) {
			continue
		}
		height := utils.IntParseRune(runes[yy][xx])
		if height-currentHeight == 1 {
			ends := getHickingTrailsEnds(runes, xx, yy)
			result = append(result, ends...)
		}
	}
	return result
}

func inRange(runes [][]rune, x, y int) bool {
	return y >= 0 && y < len(runes) && x >= 0 && x < len(runes[0])
}

func uniqueElements(all []point) int {
	pMap := make(map[point]bool, len(all))
	for _, p := range all {
		log.Printf("pMap %#v contains %v %v", pMap, p, pMap[p])
		pMap[p] = true
	}
	return len(pMap)
}
