package main

import (
	"fmt"
	"time"
)

type WaitGroup struct {
	sem   chan struct{}
	count int
	zero  chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem:  make(chan struct{}, 1),
		zero: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.sem <- struct{}{}
	defer func() { <-wg.sem }()

	if wg.count == 0 && delta > 0 {
		wg.zero = make(chan struct{})
	}

	wg.count += delta
	if wg.count < 0 {
		panic("WaitGroup: counter is negative")
	}
}

func (wg *WaitGroup) Done() {
	wg.sem <- struct{}{}
	defer func() { <-wg.sem }()
	wg.count--

	if wg.count < 0 {
		panic("WaitGroup: counter is negative")
	}

	if wg.count == 0 {
		close(wg.zero)
	}
}

func (wg *WaitGroup) Wait() {
	wg.sem <- struct{}{}
	zero := wg.zero
	done := wg.count == 0
	<-wg.sem

	if !done {
		<-zero
	}
}

func main() {
	wg := NewWaitGroup()
	results := make([]int, 3)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(100-id*20) * time.Millisecond)
			results[id] = id * id
			fmt.Printf("горутина %d завершилась, результат=%d\n", id, results[id])
		}(i)
	}

	wg.Wait()
	fmt.Println("все горутины завершены:", results)
}
