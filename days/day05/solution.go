package day05

import (
	"adventofcode2024/pkg/utils"
	"log"
	"slices"
)

func Solve(input []string) {
	rules, numbers := parseInput(input)

	log.Printf("day 5 part 1: %d", part1(rules, numbers))
	log.Printf("day 5 part 2: %d", part2(rules, numbers))
}

func parseInput(input []string) (map[int]map[int]bool, [][]int) {
	rules := make(map[int]map[int]bool)
	numbers := make([][]int, 0)

	isReadingRules := true
	for _, line := range input {
		if len(line) == 0 {
			isReadingRules = false
		} else if isReadingRules {
			arr := utils.StringToIntArray(line, "|")
			entries := rules[arr[0]]
			if entries == nil {
				entries = make(map[int]bool)
			}
			entries[arr[1]] = true
			rules[arr[0]] = entries
		} else {
			numbers = append(numbers, utils.StringToIntArray(line, ","))
		}
	}
	return rules, numbers
}

func part1(rules map[int]map[int]bool, numbers [][]int) int {
	var result int
	for _, nums := range numbers {
		if isCorrectOrder(nums, rules) {
			result += nums[len(nums)/2]
		}
	}
	return result
}

func part2(rules map[int]map[int]bool, numbers [][]int) int {
	var result int

	for _, nums := range numbers {
		if !isCorrectOrder(nums, rules) {
			correctOrder(&nums, rules)
			result += nums[len(nums)/2]
		}
	}
	return result
}

func correctOrder(numbers *[]int, rules map[int]map[int]bool) {
	slices.SortFunc(*numbers, func(a, b int) int {
		r := rules[a]
		if r == nil || !r[b] {
			return 1
		} else {
			return -1
		}
	})
}

func isCorrectOrder(numbers []int, rules map[int]map[int]bool) bool {
	for i, n := range numbers {
		for _, m := range numbers[i+1:] {
			r := rules[n]
			if r == nil || !r[m] {
				return false
			}
		}
	}
	return true
}
