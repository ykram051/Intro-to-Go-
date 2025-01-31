package main

import "fmt"

func v1(){
	
	var grade1 , grade2 , grade3 , grade4 , grade5 float64
	fmt.Println("enter grade 1 : ")
	fmt.Scan(&grade1)
	fmt.Println("enter grade 2 : ")
	fmt.Scan(&grade2)
	fmt.Println("enter grade 3 : ")
	fmt.Scan(&grade3)
	fmt.Println("enter grade 4 : ")
	fmt.Scan(&grade4)
	fmt.Println("enter grade 5 : ")
	fmt.Scan(&grade5)

	fmt.Printf("The average is : %f ", average(grade1,grade2,grade3,grade4,grade5))

}

func average(grade1 , grade2 , grade3 , grade4 , grade5 float64 ) float64 {
	return (grade1 + grade2 + grade3 + grade4 + grade5)/5

}