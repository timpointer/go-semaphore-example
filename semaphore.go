package go_semaphore_example

type Semaphore interface {
	Signal()
	MultiSignal(n int)
	Wait()
}

const signal = 1
const size = 10000

func NewSemaphore(v int) Semaphore {
	s := &semaphore{
		ch: make(chan int, size),
	}
	for i := 0; i < v; i++ {
		s.ch <- signal
	}
	return s
}

type semaphore struct {
	ch chan int
}

func (s *semaphore) MultiSignal(n int) {
	for i := 0; i < n; i++ {
		s.Signal()
	}
}

func (s *semaphore) Signal() {
	s.ch <- signal
}

func (s *semaphore) Wait() {
	_ = <-s.ch
}
