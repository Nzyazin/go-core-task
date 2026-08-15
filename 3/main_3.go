package main

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

}
