package main

import "math/rand"

func RandomNumbers(count int, ch chan<- int) {
	defer close(ch)
	for _ = range count {
		ch <- rand.Int()
	}
}

func main() {
	ch := make(chan int)
	go RandomNumbers(3, ch)
}
