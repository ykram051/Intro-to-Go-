package main

import ("log" 
		"um6p.ma/hello/mathutils"
	)
func f3(){
	result := mathutils.Add(1,2) // Use a function from the mathutils package
	log.Println("Result from mathutils:", result)
	log.Println("hello world")
}