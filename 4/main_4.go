package main

import "fmt"

func compareSlice(firstSlice []string, secondSlice []string) []string {
	result := make([]string, 0)
	set := make(map[string]bool)

	for _, v := range secondSlice {
		set[v] = true
	}

	for _, v := range firstSlice {
		if !set[v] {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(compareSlice(slice1, slice2))
}
