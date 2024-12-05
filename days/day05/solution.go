package day05

import (
	"adventofcode2024/pkg/utils"
	"log"
)

func Solve(input []string) {
	log.Printf("day 5 part 1: %d", part1(input))

}

func part1(input []string) int {
	var result int
	rules := make(map[int]map[int]bool, 0)

	isReadingRules := true
	for _, line := range input {
		if len(line) == 0 {
			isReadingRules = false
			continue
		}

		if isReadingRules {
			arr := utils.StringToIntArray(line, "|")
			entries := rules[arr[0]]
			if entries == nil {
				entries = make(map[int]bool)
			}
			entries[arr[1]] = true
			rules[arr[0]] = entries
		} else {
			numbers := utils.StringToIntArray(line, ",")
			if isCorrectOrder(numbers, rules) {
				result += numbers[len(numbers)/2]
			}
		}
	}
	return result
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
