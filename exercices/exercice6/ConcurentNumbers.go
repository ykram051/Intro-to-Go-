package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func square1(n int) {
	fmt.Printf("Square of %d is %d\n", n, n*n)
	defer wg.Done()
}

func main1() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go square1(i)
	}
	wg.Wait()
}
