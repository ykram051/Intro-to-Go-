package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Salary    int
	Education string
}

type People struct {
	People []Person `json:"People"`
}

type Stats struct {
	AverageAge                       float64
	AverageSalary                    float64
	Nameoftheyoungestperson          []string
	NameOfTheOldestPerson            []string
	NameOfThePersonWithHighestSalary []string
	NameOfThePersonWithLowestSalary  []string
	CountsOfEducationLevel           map[string]int
}

func main() {
	file, err := os.ReadFile("./person.json")
	if err != nil {
		log.Fatal(err)
	}

	var people People

	err = json.Unmarshal(file, &people)
	if err != nil {
		log.Fatal(err)
	}
	stats := Stats{AverageAge: averageage(people),
		Nameoftheyoungestperson:          nameoftheyoungestperson(people),
		NameOfTheOldestPerson:            nameOfTheOldestPerson(people),
		NameOfThePersonWithHighestSalary: nameOfThePersonWithHighestSalary(people),
		NameOfThePersonWithLowestSalary:  nameOfThePersonWithLowestSalary(people),
		AverageSalary:                    averageSalary(people),
		CountsOfEducationLevel:           countsOfEducationLevel(people),
	}
	ofile, _ := json.Marshal(stats)
	ioutil.WriteFile("output.json", ofile, 0644)

}

func averageage(listofpersons People) float64 {
	var average float64 = 0
	length := len(listofpersons.People)
	for i := 0; i < length; i++ {
		average += float64(listofpersons.People[i].Age)

	}
	return average / float64(length)
}

func nameoftheyoungestperson(listofpersons People) []string {
	var minAge int = 5000
	var temp []string
	length := len(listofpersons.People)

	for i := 0; i < length; i++ {
		if listofpersons.People[i].Age < minAge {
			minAge = listofpersons.People[i].Age
			temp = []string{listofpersons.People[i].Name}
		} else if listofpersons.People[i].Age == minAge {
			temp = append(temp, listofpersons.People[i].Name)
		}
	}
	return temp
}

func nameOfTheOldestPerson(listofpersons People) []string {
	var maxAge int
	var temp []string
	length := len(listofpersons.People)

	for i := 0; i < length; i++ {
		if listofpersons.People[i].Age > maxAge {
			maxAge = listofpersons.People[i].Age
			temp = []string{listofpersons.People[i].Name}
		} else if listofpersons.People[i].Age == maxAge {
			temp = append(temp, listofpersons.People[i].Name)
		}
	}
	return temp
}

func averageSalary(listofpersons People) float64 {
	var average float64 = 0
	length := len(listofpersons.People)
	for i := 0; i < length; i++ {
		average += float64(listofpersons.People[i].Salary)

	}
	return average / float64(length)
}

func nameOfThePersonWithLowestSalary(listofpersons People) []string {
	var minSalary int = 10e16
	var temp []string
	length := len(listofpersons.People)

	for i := 0; i < length; i++ {
		if listofpersons.People[i].Salary < minSalary {
			minSalary = listofpersons.People[i].Salary
			temp = []string{listofpersons.People[i].Name}
		} else if listofpersons.People[i].Salary == minSalary {
			temp = append(temp, listofpersons.People[i].Name)
		}
	}
	return temp
}

func nameOfThePersonWithHighestSalary(listofpersons People) []string {
	var maxsalary int
	var temp []string
	length := len(listofpersons.People)

	for i := 0; i < length; i++ {
		if listofpersons.People[i].Salary > maxsalary {
			maxsalary = listofpersons.People[i].Salary
			temp = []string{listofpersons.People[i].Name}
		} else if listofpersons.People[i].Salary == maxsalary {
			temp = append(temp, listofpersons.People[i].Name)
		}
	}
	return temp
}

func countsOfEducationLevel(listofpersons People) map[string]int {
	length := len(listofpersons.People)
	m := make(map[string]int)
	for i := 0; i < length; i++ {
		_, ok := m[listofpersons.People[i].Education]
		if ok {
			m[listofpersons.People[i].Education]++
		} else {
			m[listofpersons.People[i].Education] = 0
		}
	}
	return m
}
