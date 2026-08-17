package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()

		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3

		close(ch1)
		close(ch2)
		close(ch3)
	}()

	res := mergeChannels(ch1, ch2, ch3)

	for v := range res {
		fmt.Println(v)
	}
}
