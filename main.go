package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"sync"
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

// AreaData represents the state of each area
type AreaData struct {
	Area            string
	FlowRate        float64
	Leakage         bool
	NormalFlowRate  float64
	FairPercentage  float64
}

// Global variable to store the state of each area
var areas []AreaData
var areasMutex sync.Mutex

// initializeAreas initializes the state of each area
func initializeAreas() {
	areas = []AreaData{
		{"Kondele Area", 0, false, rand.Float64() * 100, 30},
		{"Manyatta Area", 0, false, rand.Float64() * 100, 40},
		{"Mamboleo Area", 0, false, rand.Float64() * 100, 30},
	}
}

// simulateAdminData simulates real-time admin sensor data
func simulateAdminData() []AdminSensorData {
	areasMutex.Lock()
	defer areasMutex.Unlock()

	for i := range areas {
		if areas[i].Leakage {
			areas[i].FlowRate = areas[i].NormalFlowRate + (rand.Float64() * 200 - 100) // Simulate leakage
		} else {
			areas[i].FlowRate = areas[i].NormalFlowRate + (rand.Float64() * 20 - 10) // Normal fluctuation
		}
		// Ensure flow rate is positive
		if areas[i].FlowRate < 0 {
			areas[i].FlowRate = 0
		}
	}

	data := make([]AdminSensorData, len(areas))
	for i, area := range areas {
		data[i] = AdminSensorData{
			Area:           area.Area,
			FlowRate:       area.FlowRate,
			Leakage:        area.Leakage,
			FairPercentage: area.FairPercentage,
		}
	}

	return data
}

// simulateUserData simulates real-time user sensor data
func simulateUserData() []UserSensorData {
	return []UserSensorData{
		{FlowRate: rand.Float64() * 100, Cost: rand.Float64() * 100},
	}
}

// checkForLeakage checks and updates the leakage status based on flow rate
func checkForLeakage() {
	areasMutex.Lock()
	defer areasMutex.Unlock()

	for i := range areas {
		flowRate := rand.Float64() * 100
		if flowRate > areas[i].NormalFlowRate*1.5 || flowRate < areas[i].NormalFlowRate*0.5 {
			areas[i].Leakage = true
		} else {
			areas[i].Leakage = false
		}
		areas[i].FlowRate = flowRate
	}
}

// simulateLeakage simulates leakage starting after 1 minute and stopping after 2 minutes
func simulateLeakage() {
	for {
		time.Sleep(1 * time.Minute) // Wait for 1 minute before starting the leakage
		areasMutex.Lock()
		for i := range areas {
			areas[i].Leakage = true
		}
		areasMutex.Unlock()

		time.Sleep(2 * time.Minute) // Wait for 2 minutes before stopping the leakage
		areasMutex.Lock()
		for i := range areas {
			areas[i].Leakage = false
		}
		areasMutex.Unlock()
	}
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin.html")
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
	tmpl, err := template.ParseFiles("templates/user.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signup.html")
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
	initializeAreas()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/admin", adminPageHandler)
	http.HandleFunc("/login", loginPageHandler)

	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/user", userPageHandler)
	http.HandleFunc("/admin-data", adminDataHandler)
	http.HandleFunc("/user-data", userDataHandler)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Periodically check for leakage every 2 minutes
	go func() {
		for range time.Tick(2 * time.Minute) {
			checkForLeakage()
		}
	}()

	// Simulate leakage starting after 1 minute and stopping after 2 minutes
	go simulateLeakage()

	http.ListenAndServe(":8060", nil)
}
