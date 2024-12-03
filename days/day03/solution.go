package day03

import (
	"adventofcode2024/pkg/utils"
	"log"
	"regexp"
)

func Solve(input []string) {
	log.Printf("day 3 part 1: %d", part1(input))
	log.Printf("day 3 part 2: %d", part2(input))
}

func part1(input []string) int {
	var result int

	muls, err := findMuls(input, true)
	if err != nil {
		log.Fatalf("can't create regexp: %v\n", err)
	}

	nums, err := mapMuls(muls)
	if err != nil {
		log.Fatalf("can't map muls: %v\n", err)
	}

	for _, arr := range nums {
		result += arr[0] * arr[1]
	}

	return result
}

// 97836217 - to hight
func part2(input []string) int {
	var result int

	muls, err := findMuls(input, false)
	if err != nil {
		log.Fatalf("can't create regexp: %v\n", err)
	}

	nums, err := mapMuls(muls)
	if err != nil {
		log.Fatalf("can't map muls: %v\n", err)
	}

	for _, arr := range nums {
		result += arr[0] * arr[1]
	}

	return result
}

func findMuls(input []string, alwaysEnable bool) ([]string, error) {
	r, err := regexp.Compile("don't\\(\\)|do\\(\\)|mul\\([\\d+]*[\\,][\\d+]*\\)")
	if err != nil {
		return nil, err
	}

	isEnabled := true
	muls := make([]string, 0)
	for _, line := range input {

		all := r.FindAllString(line, -1)
		for _, match := range all {
			if match == "don't()" {
				isEnabled = false
			} else if match == "do()" {
				isEnabled = true
			} else if alwaysEnable || isEnabled {
				muls = append(muls, match)
			}
		}
	}

	return muls, nil
}

func mapMuls(muls []string) ([][]int, error) {
	r, err := regexp.Compile("[\\d+]*[\\,][\\d+]*")
	if err != nil {
		return nil, err
	}

	result := make([][]int, len(muls))

	for i, line := range muls {
		str := r.FindString(line)
		result[i] = utils.StringToIntArray(str, ",")
	}

	return result, nil
}
