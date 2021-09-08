package go_semaphore_example

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Example_3_1_Signaling() {
	sem := NewSemaphore(0)
	go func() {
		sem.Wait()
		Statement("b1")
	}()

	go func() {
		Statement("a1")
		sem.Signal()
	}()

	time.Sleep(time.Second)
	// Output:
	// statement a1
	// statement b1
}

func Example_3_3_Rendezvous() {
	aArrived := NewSemaphore(0)
	bArrived := NewSemaphore(0)

	go func() {
		Statement("b1")
		bArrived.Signal()
		aArrived.Wait()
		Statement("b2")
	}()

	go func() {
		Statement("a1")
		bArrived.Wait()
		aArrived.Signal()
		Statement("a2")
	}()

	time.Sleep(time.Second)
	// Output:
	// statement a1
	// statement b1
	// statement a2
	// statement b2
}

func Example_3_4_Mutual() {
	count := 0
	locker := NewLocker()

	for i := 0; i < 1000; i++ {
		go func() {
			locker.Lock()
			count = count + 1
			locker.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
	// Output:
	// 1000
}

func Test_3_5_Multiplex(t *testing.T) {
	// five tokens
	tokens := NewSemaphore(5)

	for i := 0; i < 1000; i++ {
		// only 5 goroutines can run concurrency
		go func(num int) {
			tokens.Wait()
			fmt.Printf("%v num %d\n", time.Now().Format(time.Stamp), num)
			time.Sleep(time.Second)
			tokens.Signal()
		}(i)
	}

	time.Sleep(time.Hour)
}

func Test_3_6_Barrier_implement_with_waitGroup(t *testing.T) {
	n := 5
	group := &sync.WaitGroup{}
	group.Add(n)

	for i := 0; i < n; i++ {
		go func(num int) {
			fmt.Printf("%v num %d\n", now(), num)
			time.Sleep(time.Second)
			group.Done()
			group.Wait()
			fmt.Printf("%v num %d critical point\n", now(), num)
		}(i)
	}

	time.Sleep(time.Hour)
}

// compare with waitGroup implementation,which is simpler?
func Test_3_6_Barrier(t *testing.T) {
	n := 5
	count := 0
	locker := NewLocker()
	barrier := NewSemaphore(0)

	for i := 0; i < n; i++ {
		go func(num int) {
			fmt.Printf("%v num %d\n", now(), num)
			sleep(1)

			locker.Lock()
			count += 1
			if count == n {
				barrier.Signal()
			}
			locker.Unlock()
			barrier.Wait()
			barrier.Signal()

			fmt.Printf("%v num %d critical point\n", now(), num)
		}(i)
	}

	time.Sleep(time.Hour)
}

func Test_3_7_Reusable_Barrier(t *testing.T) {
	n := 5
	count := 0
	locker := NewLocker()
	turnstile := NewSemaphore(0)
	turnstile2 := NewSemaphore(1) // hit: init one single

	for i := 0; i < n; i++ {
		go func(num int) {
			for {
				fmt.Printf("%v num %d\n", now(), num)
				sleep(1)

				locker.Lock()
				count += 1
				if count == n {
					turnstile2.Wait()
					turnstile.Signal()
				}
				locker.Unlock()
				turnstile.Wait()
				turnstile.Signal()

				fmt.Printf("%v num %d critical point\n", now(), num)

				locker.Lock()
				count -= 1
				if count == 0 {
					turnstile.Wait()
					turnstile2.Signal()
				}
				locker.Unlock()
				turnstile2.Wait()
				turnstile2.Signal()

			}
		}(i)
	}

	time.Sleep(time.Hour)
}

func Test_3_7_6_Preload_turnstile(t *testing.T) {
	n := 5
	count := 0
	locker := NewLocker()
	turnstile := NewSemaphore(0)
	turnstile2 := NewSemaphore(0)

	for i := 0; i < n; i++ {
		go func(num int) {
			for {
				fmt.Printf("%v num %d\n", now(), num)
				sleep(1)

				locker.Lock()
				count += 1
				if count == n {
					turnstile.MultiSignal(n)
				}
				locker.Unlock()
				turnstile.Wait()

				fmt.Printf("%v num %d critical point\n", now(), num)

				locker.Lock()
				count -= 1
				if count == 0 {
					turnstile2.MultiSignal(n)
				}
				locker.Unlock()
				turnstile2.Wait()

			}
		}(i)
	}

	time.Sleep(time.Hour)
}

func now() string {
	return time.Now().Format(time.Stamp)
}

func sleep(second int) {
	time.Sleep(time.Duration(second) * time.Second)
}

func Statement(i string) {
	fmt.Printf("statement %s\n", i)
}
