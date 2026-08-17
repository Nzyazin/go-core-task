package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name   string
		inputs [][]int
		want   []int
	}{
		{"three channels", [][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
		{"no channels", [][]int{}, []int{}},
		{"all empty", [][]int{{}, {}, {}}, []int{}},
		{"empty and non-empty", [][]int{{}, {7, 8}}, []int{7, 8}},
		{"duplicates", [][]int{{1}, {1}, {1}}, []int{1, 1, 1}},
		{"different lengths", [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			channels := make([]<-chan int, 0, len(test.inputs))
			for _, data := range test.inputs {
				ch := make(chan int)
				channels = append(channels, ch)

				go func(data []int, ch chan int) {
					defer close(ch)
					for _, v := range data {
						ch <- v
					}
				}(data, ch)
			}

			got := []int{}

			for v := range mergeChannels(channels...) {
				got = append(got, v)
			}

			sort.Ints(got)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("mergeChannels = %v, want %v", got, test.want)
			}
		})
	}
}
