package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// AdminSensorData represents the data sent to the admin dashboard
type AdminSensorData struct {
	Area           string  `json:"area"`
	FlowRate       float64 `json:"flow_rate"`
	Leakage        bool    `json:"leakage"`
	FairPercentage float64 `json:"fair_percentage"`
}

// UserSensorData represents the data sent to the user dashboard
type UserSensorData struct {
	FlowRate float64 `json:"flow_rate"`
	Cost     float64 `json:"cost"`
}

// simulateAdminData simulates real-time admin sensor data
func simulateAdminData() []AdminSensorData {
	return []AdminSensorData{
		{"Kondele Area", rand.Float64() * 100, rand.Intn(2) == 1, 30},
		{"Manyatta Area", rand.Float64() * 100, rand.Intn(2) == 1, 40},
		{"Mamboleo Area", rand.Float64() * 100, rand.Intn(2) == 1, 30},
	}
}


// simulateUserData simulates real-time user sensor data
func simulateUserData() []UserSensorData {
	return []UserSensorData{
		{FlowRate: rand.Float64() * 100, Cost: rand.Float64() * 100},
	}
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/admin.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func userPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/user.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func adminDataHandler(w http.ResponseWriter, r *http.Request) {
	data := simulateAdminData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func userDataHandler(w http.ResponseWriter, r *http.Request) {
	data := simulateUserData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/admin", adminPageHandler)
	http.HandleFunc("/user", userPageHandler)
	http.HandleFunc("/admin-data", adminDataHandler)
	http.HandleFunc("/user-data", userDataHandler)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	http.ListenAndServe(":8090", nil)
}
