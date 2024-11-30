package utils

import (
	"slices"
	"testing"
)

func TestArrayUtils(t *testing.T) {
	t.Run("map strings to int", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, _ := StringsToInt([]string{"1", "2", "3"})
		if !slices.Equal(want, got) {
			t.Errorf("want\n%v\ngot\n%v", want, got)
		}
	})
}
