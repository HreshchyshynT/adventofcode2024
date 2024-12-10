package day09

import (
	"adventofcode2024/pkg/aoc"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		dayInput, _ := aoc.Get(9)
		tests := []struct {
			want  int
			input string
		}{
			{want: 60, input: "12345"},
			{want: 1928, input: "2333133121414131402"},
			{want: 6310675819476, input: dayInput[0]},
		}
		for _, test := range tests {
			want := test.want
			got := part1(test.input)
			if want != got {
				t.Errorf("want %d got %d", want, got)
			}
		}
	})
}
