package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result := sliceExample(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}

	result := addElements(input, 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3}
	copySlice := copySlice(original)

	original[0] = 999

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(copySlice, expected) {
		t.Errorf("copy was changed: expected %v, got %v", expected, copySlice)
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}

	result := removeElement(input, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestRemoveElementFirst(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2, 3}

	result := removeElement(input, 0)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestRemoveElementLast(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2}

	result := removeElement(input, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	input := []int{1, 2, 3}

	result := removeElement(input, 10)

	if !reflect.DeepEqual(result, input) {
		t.Errorf("expected %v, got %v", input, result)
	}
}
