package day01

import (
	"adventofcode2024/pkg/aoc"
	"adventofcode2024/pkg/utils"
	"fmt"
	"log"
	"slices"
)

func Solve(part int) {
	log.Printf("solving 1st day, part: %d", part)
	input, err := aoc.Get(1)
	if err != nil {
		log.Fatal("error getting input for day 1")
	}
	switch part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	}

}

func part1(input []string) {
	left := make([]int, len(input))
	right := make([]int, len(input))

	for i, line := range input {
		nums := utils.StringToIntArray(line)
		left[i] = nums[0]
		right[i] = nums[1]
	}

	cmp := func(i, j int) int {
		return i - j
	}

	slices.SortFunc(left, cmp)
	slices.SortFunc(right, cmp)

	total := 0

	for i, n := range left {
		total += absDiff(n, right[i])
	}

	fmt.Printf("part 1 solution: %d\n", total)
}

func part2(input []string) {

}

func absDiff(l, r int) int {
	result := l - r
	if result < 0 {
		result *= -1
	}
	return result
}
