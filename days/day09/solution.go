package day09

import (
	"adventofcode2024/pkg/utils"
	"log"
)

func Solve(input []string) {
	log.Printf("day 9 part 1: %d", part1(input[0]))

}

func part1(line string) int {
	runes := []rune(line)
	left, right := 0, len(runes)-1
	var result int

	empties, files := 0, 0
	leftId, rightId := 0, len(runes)/2

	expandedIndex := 0
	for left < right {
		if files == 0 && right%2 == 0 {
			files = utils.IntParseRune(runes[right])
		}
		if empties == 0 && left%2 != 0 {
			empties = utils.IntParseRune(runes[left])
		}

		if files == 0 {
			right--
		}
		if empties == 0 {
			if left%2 == 0 {
				// it is file, add to checksum
				count := utils.IntParseRune(runes[left])
				result += findSumFor(leftId, expandedIndex, count)
				leftId += 1
				expandedIndex += count
			}
			left++
		}
		if empties != 0 && files != 0 {
			wasEmpties := empties
			empties -= files
			var count int
			if empties >= 0 {
				files = 0
				count = wasEmpties - empties
			} else {
				files -= wasEmpties
				count = wasEmpties
				empties = 0
			}
			result += findSumFor(rightId, expandedIndex, count)
			expandedIndex += count
			if empties == 0 {
				left++
			}
			if files == 0 {
				right--
				rightId--
			}
		}
	}
	if files > 0 {
		result += findSumFor(rightId, expandedIndex, files)
	}

	return result
}

func findSumFor(fileId, startIndex, count int) int {
	result := (startIndex + startIndex + count - 1) * count / 2 * fileId
	return result
}
