package main

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := NewWaitGroup()

	var result int
	var mu sync.Mutex

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			result++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if result != 3 {
		t.Fatalf("result = %d, want 3", result)
	}
}
