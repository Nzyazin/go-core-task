package main

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
	wg.Add(1)
	wg.Done()
}
