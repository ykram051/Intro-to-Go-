package main

import (
	"fmt"
	"time"
)

func Say3() {
	fmt.Println("3 ")
	time.Sleep(3 * time.Second)
}

func Say2() {
	fmt.Println("2 ")
	time.Sleep(2 * time.Second)
}

func Say1() {
	fmt.Println("1 ")
	time.Sleep(1 * time.Second)

}

func v1() {
	go Say3()
	go Say2()
	go Say1()

	fmt.Println("hello")
	time.Sleep(4 * time.Second) // if we dont add this we will print only hello bedause its like the process main finishes and exits before hello is executed

}
