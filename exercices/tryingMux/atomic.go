package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Ressource struct {
	counter int32
	mux     sync.Mutex
}

func (r *Ressource) increment(value int) {
	// The operation is performed atomically, meaning no other goroutine can observe counter in an intermediate state.
	atomic.AddInt32(&r.counter, int32(value))
	// or using a mutex to protect the counter
	r.mux.Lock()
	defer r.mux.Unlock()
	r.counter++
}

func main() {
	var wg sync.WaitGroup
	r := &Ressource{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r.increment(1)
		}()
	}
	wg.Wait()
	fmt.Println(r.counter)
}