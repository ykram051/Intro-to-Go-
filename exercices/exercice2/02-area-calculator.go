// filepath: /C:/Users/USER/Desktop/Intro-to-Go-/exercices/exercice2/02-area-calculator.go
package main

import (
    "fmt"
    "example.com/exercice2/areas"
)

func main() {
    var userLength int
    fmt.Print("Enter the length: ")
    fmt.Scan(&userLength)

    var userWidth int
    fmt.Print("Enter the width: ")
    fmt.Scan(&userWidth)

    if userLength == 0 {
        fmt.Println("Length is null")
    }
    if userWidth == 0 {
        fmt.Println("Width is null")
    }

    // Convert int to float64
    length := float64(userLength)
    width := float64(userWidth)

    area := areas.CalculateArea(length, width)
    fmt.Printf("The area using the function: %f\n", area)

    if isLengthEmpty := userLength == 0; isLengthEmpty {
        fmt.Println("Empty")
    }// just trying this notation
}