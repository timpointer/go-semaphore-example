package go_semaphore_example

type Locker interface {
	Lock()
	Unlock()
}

func NewLocker() Locker {
	return &locker{
		NewSemaphore(1),
	}
}

type locker struct {
	sem Semaphore
}

func (l *locker) Lock() {
	l.sem.Wait()
}

func (l *locker) Unlock() {
	l.sem.Signal()
}
