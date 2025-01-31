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
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}
type Users struct {
	Userlist []User `json:"users"`
}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.ReadFile("users.json")
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		log.Printf("Error reading users file: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func handleUsersusingpath(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "User ID is missing in the path", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(parts[2]) // Convert the userID from the path to an integer
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var users Users
	file, err := os.ReadFile("./users.json")
	if err != nil {
		http.Error(w, "error reading the json file", http.StatusInternalServerError)
	}
	err = json.Unmarshal(file, &users)
	if err != nil {
		http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
	}
	var result Users
	result.Userlist = []User{}
	for i, _ := range users.Userlist {
		if users.Userlist[i].Id == userID {
			result.Userlist = append(result.Userlist, users.Userlist[i])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Failed to encode user data to JSON", http.StatusInternalServerError)
		return
	}
}

func handleGetUsersByQuery(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("id")
	if r.Method != http.MethodGet {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)

	}
	ids := strings.Split(path, ",")
	if len(ids) == 0 {
		http.Error(w, "empty string", http.StatusBadRequest)
	}

	file, err := os.ReadFile("./users.json")
	if err != nil {
		http.Error(w, "error reading json", http.StatusBadRequest)
	}
	var users Users
	err = json.Unmarshal(file, &users)
	if err != nil {
		http.Error(w, "error unpacking json", http.StatusBadRequest)
	}
	var result Users
	for i, _ := range ids {
		for j, _ := range users.Userlist {
			val, err := strconv.Atoi(ids[i])
			if err != nil {
				http.Error(w, "Invalid user ID in query", http.StatusBadRequest)
				return
			}
			if users.Userlist[j].Id == val {

				result.Userlist = append(result.Userlist, users.Userlist[j])

			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "Failed to encode user data to JSON", http.StatusInternalServerError)
		return
	}
}

func handleByQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the users JSON file
	file, err := os.ReadFile("users.json")
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		log.Printf("Error reading users file: %v", err)
		return
	}

	var allUsers Users
	if err := json.Unmarshal(file, &allUsers); err != nil {
		http.Error(w, "Error parsing users JSON", http.StatusInternalServerError)
		log.Printf("Error unmarshaling users data: %v", err)
		return
	}

	// Filter users based on query parameters
	query := r.URL.Query()
	ids := query["id"]
	names := query["name"]

	var filteredUsers Users
	for _, user := range allUsers.Userlist {
		idMatch := false
		nameMatch := false

		if len(ids) > 0 { // Check if ID parameter is present
			for _, idStr := range ids {
				id, err := strconv.Atoi(idStr)
				if err == nil && user.Id == id {
					idMatch = true
					break
				}
			}
		} else {
			idMatch = true // If no ID filter is present, don't filter by ID
		}

		if len(names) > 0 { // Check if Name parameter is present
			for _, name := range names {
				if strings.EqualFold(name, user.Name) {
					nameMatch = true
					break
				}
			}
		} else {
			nameMatch = true // If no Name filter is present, don't filter by Name
		}

		if idMatch && nameMatch {
			filteredUsers.Userlist = append(filteredUsers.Userlist, user)
		}
	}

	// Set response header and encode result to JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredUsers); err != nil {
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
		log.Printf("Error encoding users data: %v", err)
	}
}
func addusertojson(w http.ResponseWriter, filepath string, newuser User) {
	file, err := os.ReadFile("./users.json")
	if err != nil {
		http.Error(w, "error reading the json file", http.StatusInternalServerError)
	}
	var users Users
	err = json.Unmarshal(file, &users)
	if err != nil {
		http.Error(w, "error unpacking the json file", http.StatusInternalServerError)
	}
	users.Userlist = append(users.Userlist, newuser)
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully"})

}
func adduser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newuser User
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		http.Error(w, "not valid user", http.StatusBadRequest)
		return
	}
	addusertojson(w, "./users.json", newuser)

}
func main() {
	//http.HandleFunc("/users", handleGetAllUsers)
	http.HandleFunc("/users/", handleUsersusingpath)
	http.HandleFunc("/users", handleGetUsersByQuery)
	http.HandleFunc("/users/add", adduser)
	log.Println("Server starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
