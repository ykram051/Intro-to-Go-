package exercice3

import "fmt"

func v2() {

	var grade1, grade2, grade3, grade4, grade5 float64
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
	aver := average2(grade1, grade2, grade3, grade4, grade5)
	fmt.Printf("The average is : %f \n", aver)
	fmt.Printf("The letter corresponding  is : ")
	switch {
	case aver >= 90:
		fmt.Println("A")
	case aver >= 80:
		fmt.Println("B")
	case aver >= 70:
		fmt.Println("C")
	case aver >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

}

func average2(grade1, grade2, grade3, grade4, grade5 float64) float64 {
	return (grade1 + grade2 + grade3 + grade4 + grade5) / 5

}
