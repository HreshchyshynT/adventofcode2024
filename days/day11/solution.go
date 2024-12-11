package day11

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve(input []string) {
	input = strings.Split(input[0], " ")
	log.Printf("day 11 part 1:%v %d", input, part1(input, 25))
	log.Printf("day 11 part 2:%v %d", input, part1(input, 75))

}

func part1(stones []string, blinks int) int {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

var cache = make(map[string][]string)

func blink(stones []string) []string {
	result := make([]string, 0, len(stones)*2)
	for _, n := range stones {
		c := cache[n]
		if c == nil {
			c = change(n)
			cache[n] = c
		}
		result = append(result, c...)
	}
	return result
}

func change(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		str := fmt.Sprint(stone)
		left, right := str[:len(str)/2], str[len(str)/2:]
		right = removeZeroes(right)
		left = removeZeroes(left)
		return []string{left, right}

	}
	n, _ := strconv.Atoi(stone)
	return []string{fmt.Sprint(n * 2024)}
}

func removeZeroes(str string) string {
	for i, r := range str {
		if r != '0' {
			return str[i:]
		}
	}
	return "0"
}
