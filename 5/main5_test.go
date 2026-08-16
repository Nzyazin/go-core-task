package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name      string
		first     []int
		second    []int
		want      []int
		wantFound bool
	}{
		{"example1", []int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, []int{3, 64}, true},
		{"example2", []int{1, 2}, []int{3, 4}, []int{}, false},
		{"example3", []int{}, []int{}, []int{}, false},
		{"example4", []int{1, 2}, []int{}, []int{}, false},
		{"example5", []int{}, []int{1, 2}, []int{}, false},
		{"example6", []int{1, 2}, []int{2, 2, 1, 1}, []int{1, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists, got := findIntersection(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) || exists != tt.wantFound {
				t.Errorf("findIntersection = %v, %v, want %v, %v", exists, got, tt.want, tt.wantFound)
			}
		})
	}
}
