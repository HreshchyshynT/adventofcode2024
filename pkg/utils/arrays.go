package utils

import "strconv"

func StringsToInt(strings []string) ([]int, error) {
	result := make([]int, len(strings))
	for i, str := range strings {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}
