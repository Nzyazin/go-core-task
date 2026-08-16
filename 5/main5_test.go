package main

import "testing"

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{"example1", []int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, []int{3, 64}},
		{"example2", []int{1, 2}, []int{3, 4}, []int{}},
		{"example3", []int{}, []int{}, []int{}},
		{"example4", []int{1, 2}, []int{}, []int{}},
		{"example5", []int{}, []int{1, 2}, []int{}},
		{"example5", []int{1, 2}, []int{2, 2, 1, 1}, []int{1, 2}},
	}

	for tt := range test {
	}
}
