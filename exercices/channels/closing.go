package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("we sent here %d",i)
			ch <- i
		}
		close(ch)
	}()
	for {

		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("received ", v)
	}
}
