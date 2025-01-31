package main

import (
	"fmt"
)

func square(n int) int {
	return n * n
}

func v1() {

	v := 1
	ch := make(chan int)
	go func() {
		s := square(v)
		ch <- s

	}()

	vSquare := <-ch
	fmt.Println(v, "", vSquare)

}
