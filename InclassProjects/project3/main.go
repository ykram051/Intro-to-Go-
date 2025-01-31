package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Address struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postal_code string `json:"postal_code"`
}
type Course struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Credit int    `json:"credit"`
}

type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Major   string   `json:"major"`
	Address Address  `json:"address"`
	Courses []Course `json:"courses"`
}

type Students struct {
	StudentList []Student `json:"Students"`
}

func handleGetandAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		file, err := os.ReadFile("./students.json")
		if err != nil {
			http.Error(w, "error opening the json file ", http.StatusBadGateway)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(file)
	} else if r.Method == http.MethodPost {

		var newStudent Student
		json.NewDecoder(r.Body).Decode(&newStudent)
		if newStudent.Id != 0 {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		addStudentToJson(w, newStudent, "./students.json")

	} else {
		http.Error(w, "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}

}

func addStudentToJson(w http.ResponseWriter, newStud Student, filepath string) {
	if newStud.Name == "" || newStud.Age == 0 || newStud.Major == "" || newStud.Address.Street == "" || newStud.Address.City == "" || newStud.Address.State == "" || newStud.Address.Postal_code == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			http.Error(w, "error creating the json file", http.StatusInternalServerError)
			return
		}
		file.Close()
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		http.Error(w, "error reading the json file", http.StatusInternalServerError)
		return
	}

	if len(file) == 0 {
		file = []byte(`{"Students": []}`)
	}

	var originalStudents Students
	err = json.Unmarshal(file, &originalStudents)
	if err != nil {
		http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
		return
	}
	maxID := -1
	for _, student := range originalStudents.StudentList {
		if student.Id > maxID {
			maxID = student.Id
		}
	}
	newStud.Id = maxID + 1

	originalStudents.StudentList = append(originalStudents.StudentList, newStud)
	data, err := json.Marshal(originalStudents)
	if err != nil {
		http.Error(w, "Failed to encode Student data to JSON", http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		http.Error(w, "Failed to write to JSON file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Student added successfully"})
}

func handlebyId(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 3 {
			http.Error(w, "Student ID is missing in the path", http.StatusBadRequest)
			return
		}

		studID, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid Student ID", http.StatusBadRequest)
			return
		}

		var students Students
		file, err := os.ReadFile("./students.json")
		if err != nil {
			http.Error(w, "error reading the json file", http.StatusInternalServerError)
		}
		err = json.Unmarshal(file, &students)
		if err != nil {
			http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
		}
		var result Students
		result.StudentList = []Student{}
		found := false
		for i := range students.StudentList {
			if students.StudentList[i].Id == studID {
				result.StudentList = append(result.StudentList, students.StudentList[i])
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			http.Error(w, "Failed to encode Student data to JSON", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPut {
		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 3 {
			http.Error(w, "Student ID is missing in the path", http.StatusBadRequest)
			return
		}

		studID, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid Student ID", http.StatusBadRequest)
			return
		}

		var students Students
		file, err := os.ReadFile("./students.json")
		if err != nil {
			http.Error(w, "error reading the json file", http.StatusInternalServerError)
		}
		err = json.Unmarshal(file, &students)
		if err != nil {
			http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
		}

		var newStudent Student
		json.NewDecoder(r.Body).Decode(&newStudent)

		found := false
		for i := range students.StudentList {
			if students.StudentList[i].Id == studID {
				students.StudentList[i] = newStudent
				students.StudentList[i].Id = studID
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		data, err := json.Marshal(students)
		if err != nil {
			http.Error(w, "Failed to encode Student data to JSON", http.StatusInternalServerError)
			return
		}

		err = ioutil.WriteFile("./students.json", data, 0644)
		if err != nil {
			http.Error(w, "Failed to write to JSON file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Student updated successfully"})

	} else if r.Method == http.MethodDelete {

		path := r.URL.Path
		parts := strings.Split(path, "/")

		if len(parts) < 3 {
			http.Error(w, "Student ID is missing in the path", http.StatusBadRequest)
			return
		}

		studID, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid Student ID", http.StatusBadRequest)
			return
		}

		var students Students
		file, err := os.ReadFile("./students.json")
		if err != nil {
			http.Error(w, "error reading the json file", http.StatusInternalServerError)
		}
		err = json.Unmarshal(file, &students)
		if err != nil {
			http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
		}

		found := false
		for i, _ := range students.StudentList {
			if students.StudentList[i].Id == studID {
				students.StudentList = append(students.StudentList[:i], students.StudentList[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}

		data, err := json.Marshal(students)
		if err != nil {
			http.Error(w, "Failed to encode Student data to JSON", http.StatusInternalServerError)
			return
		}

		err = ioutil.WriteFile("./students.json", data, 0644)
		if err != nil {
			http.Error(w, "Failed to write to JSON file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
	} else {
		http.Error(w, "Only GET, PUT, and DELETE methods are allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {

	http.HandleFunc("/students", handleGetandAdd)
	http.HandleFunc("/students/", handlebyId)

	log.Println("Server starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
