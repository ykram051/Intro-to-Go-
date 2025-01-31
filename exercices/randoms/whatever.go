package main

import "fmt"

func f1() {
	var l float64
	var w float64
	fmt.Println("enter length : ")
	fmt.Scan(&l)
	fmt.Println("enter Width : ")

	fmt.Scan(&w)
	fmt.Println(calc_area(l, w))

	calc := func(l, w float64) float64 {
		return l * w
	}
	fmt.Println(calc(5, 10))
}

func calc_area(l, w float64) (float64, error) {
	if l == 0 || w == 0 {
		return 0, fmt.Errorf("length or width cannot be zero")
	} else if l < 0 || w < 0 {
		return 0, fmt.Errorf("length or width cannot be negative ")
	}

	return l * w, nil
}

func sum(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {

		total += numbers[i]
	}
	return total
}
