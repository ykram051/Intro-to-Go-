package main

import "fmt"

func main() {
	var total float64 = 0
	var num float64 = 0
	var tempgrade float64

	for tempgrade != -1 {
		total += tempgrade
		num += 1
		fmt.Scan(&tempgrade)

	}
	var count int 
	var grade float64
	for count=1 ,grade>=0,i++{
		fmt.Println("enter grade ",count,": ")
		fmt.Scanln(&grade)
		if(grade<0 || grade>100){
	fmt.Println("invalid grade")
	count--
	continue
	}
		sum+=grade
}
	}




	aver := average3(total, num)
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

func average3(total, num float64) float64 {
	return total / num

}
