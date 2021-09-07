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

func Statement(i string) {
	fmt.Printf("statement %s\n", i)
}
