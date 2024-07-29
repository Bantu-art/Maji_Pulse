package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct representing a user in the system
type User struct {
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Location     string  `json:"location,omitempty"`
	SensorID     string  `json:"sensor_id,omitempty"`     // Sensor ID associated with the user
	WaterUsage   float64 `json:"water_usage,omitempty"`   // Total water usage
	LeakDetected bool    `json:"leak_detected,omitempty"` // Indicates if a leak is detected
}

// In-memory user store
var users []User

// GetUser retrieves a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

// UpdateUser updates an existing user's information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			var updatedUser User
			if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedUser.ID = user.ID // Keep the same ID
			users[index] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// TrackWaterUsage updates the water usage for a user
func TrackWaterUsage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			var usageData struct {
				WaterUsage   float64 `json:"water_usage"`
				LeakDetected bool    `json:"leak_detected"`
			}
			if err := json.NewDecoder(r.Body).Decode(&usageData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			users[index].WaterUsage += usageData.WaterUsage
			users[index].LeakDetected = usageData.LeakDetected
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// Sample data
	users = append(users, User{ID: "1", Name: "Kamini kashita", Location: "Asengo", SensorID: "sensor_01"})
	users = append(users, User{ID: "2", Name: "John Kamau", Location: "Mamboleo", SensorID: "sensor_02"})

	// API endpoints
	router.HandleFunc("/users", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/track-usage", TrackWaterUsage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
