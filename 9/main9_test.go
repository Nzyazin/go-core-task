package main

import (
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for _, v := range []uint8{1, 2, 3} {
			in <- v
		}
	}()

	go pipeline(out, in)

	var got []float64
	for v := range out {
		got = append(got, v)
	}

	want := []float64{1, 8, 27}
	if len(got) != len(want) {
		t.Fatalf("получено %d значений, ожидалось %d (got = %v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestPipelineEmpty(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	close(in)
	go pipeline(out, in)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("из пустого входа получено значение в out")
		}
	case <-time.After(time.Second):
		t.Fatal("канал out не закрыт: возможен дедлок")
	}
}

func TestPipelineClosesOut(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		in <- 5
	}()

	go pipeline(out, in)

	select {
	case v, ok := <-out:
		if !ok {
			t.Fatal("канал out закрыт до получения значения")
		}
		if v != 125 {
			t.Fatalf("v = %v, want 125", v)
		}
	case <-time.After(time.Second):
		t.Fatal("значение не получено: возможен дедлок")
	}

	_, ok := <-out
	if ok {
		t.Fatal("канал out не закрыт после обработки всех чисел")
	}
}
