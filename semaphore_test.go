package go_semaphore_example

import (
	"fmt"
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

	time.Sleep(1 * time.Second)
	// Output:
	// statement a1
	// statement b1
	// statement a2
	// statement b2
}

func Statement(i string) {
	fmt.Printf("statement %s\n", i)
}
