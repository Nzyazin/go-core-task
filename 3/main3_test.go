package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sm := StringIntMap{
		data: make(map[string]int),
	}

	sm.Add("apple", 12)

	if sm.data["apple"] != 12 {
		t.Errorf("expected 10, got %d", sm.data["apple"])
	}
}

func TestRemoveKey(t *testing.T) {
	sm := StringIntMap{
		data: map[string]int{
			"apple":  10,
			"banana": 20,
		},
	}

	sm.RemoveKey("apple")

	if sm.Exists("apple") {
		t.Error("expected apple to be removed")
	}
}

func TestCopy(t *testing.T) {
	sm := StringIntMap{
		data: make(map[string]int),
	}

	sm.Add("apple", 10)
	sm.Add("banana", 20)

	copied := sm.Copy()

	if copied["apple"] != 10 {
		t.Errorf("expected apple to have value 10, got %d", copied["apple"])
	}

	if copied["banana"] != 20 {
		t.Errorf("expected banana to have value 20, got %d", copied["banana"])
	}

	copied["apple"] = 100

	if sm.data["apple"] != 10 {
		t.Error("changing copied map should not change original map")
	}
}

func TestExists(t *testing.T) {
	sm := StringIntMap{
		data: make(map[string]int),
	}

	sm.Add("apple", 10)

	if !sm.Exists("apple") {
		t.Error("expected apple to exist")
	}

	if sm.Exists("banana") {
		t.Error("expected banana not to exist")
	}
}

func TestGet(t *testing.T) {
	sm := StringIntMap{
		data: make(map[string]int),
	}

	sm.Add("apple", 10)

	value, ok := sm.Get("apple")

	if !ok {
		t.Error("expected apple to exist")
	}

	if value != 10 {
		t.Errorf("expected value 10, got %d", value)
	}

	_, ok = sm.Get("banana")

	if ok {
		t.Error("expected banana not to exist")
	}
}
