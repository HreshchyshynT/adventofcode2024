package day03

import (
	"slices"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("find muls", func(t *testing.T) {
		want := []string{
			"mul(2,42222222)",
			"mul(5,5)",
			"mul(11,8)",
			"mul(8,5)",
		}
		got, _ := findMuls("xmul(2,42222222)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))      mul ( 2 , 4 ), mul(1,2,2)")

		if !slices.Equal(want, got) {
			t.Errorf("want %#v got %#v", want, got)
		}
	})

	t.Run("map muls", func(t *testing.T) {
		want := [][]int{
			{2, 42222222},
			{5, 5},
			{11, 8},
			{8, 5},
		}
		got, _ := mapMuls([]string{
			"mul(2,42222222)",
			"mul(5,5)",
			"mul(11,8)",
			"mul(8,5)",
		})

		for i, intWants := range want {
			if !slices.Equal(intWants, got[i]) {
				t.Errorf("want %#v got %#v", intWants, got[i])
			}
		}
	})
	t.Run("part1", func(t *testing.T) {
		want := 161
		got := part1([]string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"})

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
