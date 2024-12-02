package day02

import (
	"adventofcode2024/pkg/utils"
	"log"
	"slices"
)

func Solve(input []string) {
	nums := parseInput(input)

	log.Printf("part 1 solution: %d\n", part1(nums))
	log.Printf("part 2 solution: %d\n", part2(nums))
}

func part1(input [][]int) int {
	var result int
	for _, nums := range input {
		if isValidSequence(nums) {
			result++
		}
	}
	return result
}

func part2(input [][]int) int {
	var result int
	for _, nums := range input {
		for i := range nums {
			if isValidSequence(nums, i) {
				result++
				break
			}
		}
	}
	return result
}

func isValidSequence(nums []int, exclude ...int) bool {
	var breaksSequence func(int, int) bool
	for i := 0; i < len(nums); i++ {
		if slices.Contains(exclude, i) {
			continue
		}
		prevIndex := i - 1
		if slices.Contains(exclude, prevIndex) {
			prevIndex--
		}
		if prevIndex < 0 {
			continue
		}

		if breaksSequence == nil {
			breaksSequence = isBreakingDec
			if nums[prevIndex] < nums[i] {
				breaksSequence = isBreakingInc
			}
		}

		prev := nums[prevIndex]
		current := nums[i]

		if breaksSequence(prev, current) {
			return false
		}
	}
	return true
}

func isBreakingDec(left, right int) bool {
	return right >= left || left-right > 3
}

func isBreakingInc(left, right int) bool {
	return left >= right || right-left > 3
}

func parseInput(input []string) [][]int {
	nums := make([][]int, len(input))

	for i, line := range input {
		nums[i] = utils.StringToIntArray(line)
	}

	return nums
}
