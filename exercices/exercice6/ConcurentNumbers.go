package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func square(n int) {
	fmt.Printf("Square of %d is %d\n", n, n*n)
	defer wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go square(i)
	}
	wg.Wait()
}
