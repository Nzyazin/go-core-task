package main

import (
	"reflect"
	"testing"
)

func TestCompareSlice(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{"empty", []string{}, []string{}, []string{}},
		{"example1",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			"empty second slice",
			[]string{"apple", "banana"},
			[]string{},
			[]string{"apple", "banana"},
		},
		{
			"empty first slice",
			[]string{},
			[]string{"banana", "date"},
			[]string{},
		},
		{
			"all elements removed",
			[]string{"a", "b", "c"},
			[]string{"c", "b", "a"},
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareSlice(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareSlice = %v, want %v", got, tt.want)
			}
		})
	}
}
