package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func (sm *StringIntMap) Add(key string, value int) {
	sm.data[key] = value
}

func (sm *StringIntMap) RemoveKey(key string) {
	delete(sm.data, key)
}

func (sm *StringIntMap) Copy() map[string]int {
	result := make(map[string]int)

	for key, value := range sm.data {
		result[key] = value
	}

	return result
}

func (sm *StringIntMap) Exists(key string) bool {
	_, ok := sm.data[key]

	return ok
}

func (sm *StringIntMap) Get(key string) (int, bool) {
	val, ok := sm.data[key]

	return val, ok
}

func main() {
	sm := StringIntMap{data: make(map[string]int)}

	sm.Add("one", 1)
	sm.Add("two", 2)
	fmt.Println("После Add:", sm.data)

	if val, ok := sm.Get("one"); ok {
		fmt.Println("Get(one):", val, ok)
	}

	fmt.Println("Exists(two):", sm.Exists("two"))
	fmt.Println("Exists(missing):", sm.Exists("missing"))

	copyMap := sm.Copy()
	copyMap["one"] = 999
	fmt.Println("Копия после изменения:", copyMap)
	fmt.Println("Оригинал:", sm.data)

	sm.RemoveKey("two")
	fmt.Println("После RemoveKey(two):", sm.data, "Exists:", sm.Exists("two"))
}
