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

func TestWaitGroupReuse(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(1)
	wg.Done()
	wg.Wait()

	wg.Add(1)
	done := make(chan struct{})
	go func() {
		defer wg.Done()
		close(done)
	}()
	wg.Wait()

	select {
	case <-done:
	default:
		t.Fatal("Wait вернулся до завершения горутины")
	}
}

func TestWaitGroupNoAdd(t *testing.T) {
	wg := NewWaitGroup()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait не вернулся при count == 0")
	}
}
