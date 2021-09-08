package go_semaphore_example

type Barrier interface {
	Wait()
}

func NewBarrier(n int) Barrier {
	return &barrier{
		n:          n,
		count:      0,
		locker:     NewLocker(),
		turnstile:  NewSemaphore(0),
		turnstile2: NewSemaphore(0),
	}
}

type barrier struct {
	n          int
	count      int
	locker     Locker
	turnstile  Semaphore
	turnstile2 Semaphore
}

func (b *barrier) phase1() {
	b.locker.Lock()
	b.count += 1
	if b.count == b.n {
		b.turnstile.MultiSignal(b.n)
	}
	b.locker.Unlock()
	b.turnstile.Wait()
}

func (b *barrier) phase2() {
	b.locker.Lock()
	b.count -= 1
	if b.count == 0 {
		b.turnstile2.MultiSignal(b.n)
	}
	b.locker.Unlock()
	b.turnstile2.Wait()
}

func (b *barrier) Wait() {
	b.phase1()
	b.phase2()
}
