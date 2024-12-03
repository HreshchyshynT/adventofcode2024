package utils

import (
	"strconv"
	"strings"
)

func StringToIntArray(str string, delimiter ...string) []int {
	var strs []string

	if len(delimiter) == 0 {
		strs = strings.Fields(str)
	} else {
		strs = strings.Split(str, delimiter[0])
	}
	result := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			result[i] = n
		}
	}
	return result
}
