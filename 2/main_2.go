package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(oldSlice []int) []int {
	newSLice := make([]int, 0)

	for i := range oldSlice {
		if oldSlice[i]%2 == 0 {
			newSLice = append(newSLice, oldSlice[i])
		}
	}

	return newSLice
}

func addElements(numbers []int, number int) []int {
	return append(numbers, number)
}

func copySlice(numbers []int) []int {
	newSlice := make([]int, len(numbers))
	copy(newSlice, numbers)

	return newSlice
}

func removeElement(numbers []int, index int) []int {
	if index < 0 || index >= len(numbers) {
		return numbers
	}

	newSlice := make([]int, 0, len(numbers)-1)
	fmt.Println(newSlice)
	newSlice = append(newSlice, numbers[:index]...)
	newSlice = append(newSlice, numbers[index+1:]...)

	return newSlice
}

func main() {
	rand.Seed(time.Now().UnixNano())
	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Original:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even:", evenSlice)

	appendedSlice := addElements(evenSlice, 999)
	fmt.Println("After adding 999:", appendedSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copy:", copiedSlice)

	if len(originalSlice) > 0 {
		originalSlice[0] = 777
	}

	fmt.Println("Original after change:", originalSlice)
	fmt.Println("Copy after original change:", copiedSlice)

	removedSlice := removeElement(originalSlice, 3)
	fmt.Println("Before remove:", originalSlice)
	fmt.Println("After remove:", removedSlice)
}
