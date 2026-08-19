package main

import (
	"fmt"
)

func toCube(v uint8) float64 {
	f := float64(v)
	return f * f * f
}

func pipeline(out chan<- float64, in <-chan uint8) {
	defer close(out)

	for v := range in {
		out <- toCube(v)
	}
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for _, v := range []uint8{1, 2, 3, 4, 5, 6, 7, 23, 222} {
			in <- v
		}
	}()

	go pipeline(out, in)

	for res := range out {
		fmt.Println(res)
	}
}
