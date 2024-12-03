package day03

import (
	"adventofcode2024/pkg/utils"
	"log"
	"regexp"
)

func Solve(input []string) {
	log.Printf("day 3 part 1: %d", part1(input))
}

func part1(input []string) int {
	var result int

	for _, line := range input {
		muls, err := findMuls(line)
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
	}

	return result
}

func findMuls(input string) ([]string, error) {
	r, err := regexp.Compile("mul\\([\\d+]*[\\,][\\d+]*\\)")
	if err != nil {
		return nil, err
	}

	return r.FindAllString(input, -1), nil
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
