package main

import "fmt"

func findIntersection(a, b []int) (bool, []int) {
	set := make(map[int]bool, len(b))
	result := make([]int, 0)

	for _, v := range b {
		set[v] = true
	}

	for _, v := range a {
		if set[v] {
			result = append(result, v)
			set[v] = false
		}
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(findIntersection(a, b))
}
