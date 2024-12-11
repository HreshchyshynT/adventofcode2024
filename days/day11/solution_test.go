package day11

import (
	"slices"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("test blink", func(t *testing.T) {
		want := []string{"1", "2024", "1", "0", "9", "9", "2021976"}
		got := blink([]string{"0", "1", "10", "99", "999"})

		if !slices.Equal(want, got) {
			t.Errorf("want %#v got %#v", want, got)
		}
	})
	t.Run("test blink many times", func(t *testing.T) {
		want := []string{"2097446912", "14168", "4048", "2", "0", "2", "4", "40", "48", "2024", "40", "48", "80", "96", "2", "8", "6", "7", "6", "0", "3", "2"}
		got := []string{"125", "17"}
		for i := 0; i < 6; i++ {
			got = blink(got)
		}

		if !slices.Equal(want, got) {
			t.Errorf("want:\n %#v got:\n %#v", want, got)
		}

	})

	t.Run("test remove zeroes", func(t *testing.T) {
		want := "720"
		got := removeZeroes("0720")

		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("part1", func(t *testing.T) {
		tests := []struct {
			want   int
			blinks int
			input  []string
		}{
			{3, 1, []string{"125", "17"}},
			{4, 2, []string{"125", "17"}},
			{5, 3, []string{"125", "17"}},
			{9, 4, []string{"125", "17"}},
			{13, 5, []string{"125", "17"}},
			{22, 6, []string{"125", "17"}},
			{31, 7, []string{"125", "17"}},
			{42, 8, []string{"125", "17"}},
			{68, 9, []string{"125", "17"}},
			{109, 10, []string{"125", "17"}},
			{170, 11, []string{"125", "17"}},
			{235, 12, []string{"125", "17"}},
			{342, 13, []string{"125", "17"}},
			{557, 14, []string{"125", "17"}},
			{853, 15, []string{"125", "17"}},
			{1298, 16, []string{"125", "17"}},
			{1951, 17, []string{"125", "17"}},
			{2869, 18, []string{"125", "17"}},
			{4490, 19, []string{"125", "17"}},
			{6837, 20, []string{"125", "17"}},
			{10362, 21, []string{"125", "17"}},
			{15754, 22, []string{"125", "17"}},
			{23435, 23, []string{"125", "17"}},
			{36359, 24, []string{"125", "17"}},
			{55312, 25, []string{"125", "17"}},
			{445882, 30, []string{"125", "17"}},
		}
		for _, test := range tests {
			want := test.want
			got := part1(test.input, test.blinks)

			if want != got {
				t.Errorf("want %d got %d", want, got)
			}
		}
	})
}
