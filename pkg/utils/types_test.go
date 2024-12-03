package utils

import (
	"slices"
	"testing"
)

func TestTypes(t *testing.T) {
	t.Run("string to int array", func(t *testing.T) {
		want := []int{1, 2}
		got := StringToIntArray("1      2")
		if !slices.Equal(want, got) {
			t.Errorf("want\n%v\ngot\n%v", want, got)
		}
	})
	t.Run("string to int array comma delimiter", func(t *testing.T) {
		want := []int{1, 2}
		got := StringToIntArray("1,2", ",")
		if !slices.Equal(want, got) {
			t.Errorf("want\n%v\ngot\n%v", want, got)
		}
	})
}
