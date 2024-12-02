package day02

import (
	"adventofcode2024/pkg/utils"
	"log"
)

func Solve(input []string) {
	nums := parseInput(input)

	log.Printf("part 1 solution: %d\n", part1(nums))
	log.Printf("part 2 solution: %d\n", part2(nums))
}

func part1(input [][]int) int {
	var result int
	for _, nums := range input {
		check := func(left, right int) bool {
			return right >= left || left-right > 3
		}
		if nums[0] < nums[1] {
			check = func(left, right int) bool {
				return left >= right || right-left > 3
			}
		}
		result++

		prev := nums[0]
		for _, n := range nums[1:] {
			if check(prev, n) {
				result--
				break
			}
			prev = n
		}
	}
	return result
}

func part2(_ [][]int) int {
	return 0
}

func parseInput(input []string) [][]int {
	nums := make([][]int, len(input))

	for i, line := range input {
		nums[i] = utils.StringToIntArray(line)
	}

	return nums
}
