package main

import "fmt"
type person struct{
	name string
	age int
}


func main(){
	p := person{name: "ikram", age: 20}
	fmt.Printf("name %s",p.name)
}
