package utils

import (
	"strconv"
	"strings"
)

func StringToIntArray(str string) []int {
	strs := strings.Fields(str)
	result := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			result[i] = n
		}
	}
	return result
}
