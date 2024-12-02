package day01

import (
	"adventofcode2024/pkg/utils"
	"fmt"
	"slices"
)

func Solve(input []string) {
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

	part1(left, right)
	part2(left, right)
}

func part1(left, right []int) {
	total := 0

	for i, n := range left {
		total += utils.AbsDiff(n, right[i])
	}

	fmt.Printf("part 1 solution: %d\n", total)
}

func part2(left, right []int) {
	total := 0

	for _, n := range left {
		total += n * timesInArray(n, right)
	}
	fmt.Printf("part 2 solution: %d\n", total)
}

func timesInArray(what int, where []int) int {
	times := 0
	l, r := 0, len(where)-1

	for r-l > 1 {
		mid := l + (r-l)/2
		if where[mid] < what {
			l = mid
		} else {
			r = mid
		}
	}

	for where[r] == what {
		times++
		r++
	}

	return times
}
