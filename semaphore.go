package go_semaphore_example

type Semaphore interface {
	Signal()
	Wait()
}

func NewSemaphore(i int) Semaphore {
	return &semaphore{
		ch: make(chan int, i),
	}
}

type semaphore struct {
	ch chan int
}

func (s *semaphore) Signal() {
	s.ch <- 1
}

func (s *semaphore) Wait() {
	_ = <-s.ch
}
