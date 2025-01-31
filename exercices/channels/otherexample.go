package main

import "fmt"

func v2() {

	ch := make(chan int, 5)
	ch <- 1
	fmt.Printf("send 1 \n")
	ch <- 2
	fmt.Printf("send 2 \n")
	ch <- 3
	fmt.Printf("send 3 \n")
	ch <- 4
	fmt.Printf("send 4 \n")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

}
