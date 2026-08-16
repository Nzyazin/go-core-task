package main

import (
	"testing"
)

func TestRandomNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 3},
		{"4", 4, 4},
		{"5", 5, 5},
		{"6", 6, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch := make(chan int)
			go RandomNumbers(test.input, ch)

			got := make([]int, 0, test.input)
			for v := range ch {
				got = append(got, v)
			}

			if len(got) != test.want {
				t.Errorf("got %d, want %d", len(got), test.want)
			}
		})
	}
}
