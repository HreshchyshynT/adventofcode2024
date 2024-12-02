package day02

import (
	"adventofcode2024/pkg/utils"
	"log"
)

func Solve(input []string) {
	nums := make([][]int, len(input))

	for i, line := range input {
		nums[i] = utils.StringToIntArray(line)
	}

	part1(nums)
}

func part1(input [][]int) {
	var result int
	log.Printf("part 1 solution: %d\n", result)
}
