package main

import "fmt"

func main() {
	var elems []string
	fmt.Println("insert the words of the list you want to check and rite exit if you want to stop . ")
	for {
		fmt.Println("insert a word ")
		var temp string
		fmt.Scan(&temp)
		if temp == "exit" {
			break
		}
		elems = append(elems, temp)
	}

	result := wordFrequencies(elems)
	for i, v := range result {
		fmt.Printf("the key %s occured %d times . ", i, v)
	}

}

func wordFrequencies(elems []string) map[string]int {
	result := make(map[string]int)
	for _, v := range elems {
		_, ok := result[v]
		if ok {
			result[v] += 1
		} else {
			result[v] = 1
		}

	}
	return result
}
