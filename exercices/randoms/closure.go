package main

import "fmt"

func f2() {
	bo := increment()
	fmt.Printf("%d\n",bo(3))
	fmt.Printf("%d",bo(4))

}

func increment() func(a int) int {
	i:=0
	return func(a int) int{
		fmt.Println("here")
		i+=a
		return i 
	}

}