package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"smart-water-management/internal/sensor"
)

// Templates holds the parsed template files
var Templates *template.Template

// Dummy population data
var areaPopulations = map[string]int{
	"Area1": 1000,
	"Area2": 1500,
	"Area3": 2000,
}

// LoadTemplates loads the templates from the filesystem
func LoadTemplates() {
	Templates = template.Must(template.ParseGlob("templates/*.html"))
}

// HomePage serves the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

// UserPage serves the user page

// AdminPage serves the admin page
func AdminPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"FlowRates":  getAreaFlowRates(),
		"FairDist":   getFairDistribution(),
		"LeakDetect": getLeakDetection(),
	}
	Templates.ExecuteTemplate(w, "admin.html", data)
}

// FairDistributionPage serves the fair distribution page
func FairDistributionPage(w http.ResponseWriter, r *http.Request) {
	data := getFairDistribution()
	Templates.ExecuteTemplate(w, "fair-distribution.html", data)
}

// LeakDetectionPage serves the leak detection page
func LeakDetectionPage(w http.ResponseWriter, r *http.Request) {
	data := getLeakDetection()
	Templates.ExecuteTemplate(w, "leak-detection.html", data)
}


// GetAreaFlowRates handles the API request for area flow rates
func GetAreaFlowRates(w http.ResponseWriter, r *http.Request) {
	data := getAreaFlowRates()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetFairDistribution handles the API request for fair distribution
func GetFairDistribution(w http.ResponseWriter, r *http.Request) {
	data := getFairDistribution()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetLeakDetection handles the API request for leak detection
func GetLeakDetection(w http.ResponseWriter, r *http.Request) {
	data := getLeakDetection()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// getAreaFlowRates retrieves simulated flow rates for areas
func getAreaFlowRates() []map[string]interface{} {
	var rates []map[string]interface{}
	areas := []string{"Area1", "Area2", "Area3"}
	for _, area := range areas {
		data := sensor.SimulateAreaSensorData(area)
		rates = append(rates, map[string]interface{}{
			"name":      data.Area,
			"flow_rate": data.FlowRate,
			"volume":    data.Volume,
		})
	}
	return rates
}

// getFairDistribution calculates and returns fair water distribution based on population

// getLeakDetection retrieves simulated leak detection data
func getLeakDetection() []map[string]interface{} {
	var leaks []map[string]interface{}
	areas := []string{"Area1", "Area2", "Area3"}
	for _, area := range areas {
		data := sensor.SimulateLeakSensorData(area)
		status := "No leak detected"
		if data.FlowRate > 18 || data.FlowRate < 2 { // Example threshold for leak detection
			status = "Leak detected"
		}
		leaks = append(leaks, map[string]interface{}{
			"area":   area,
			"status": status,
		})
	}
	return leaks
}

func GetSensorData(w http.ResponseWriter, r *http.Request) {
	data := sensor.GetSensorData() // Note: Ensure GetSensorData exists and is valid
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getFairDistribution() map[string]interface{} {
	var totalPopulation int
	for _, population := range areaPopulations {
		totalPopulation += population
	}

	distribution := make(map[string]interface{})
	for area, population := range areaPopulations {
		distribution[area] = float64(population) / float64(totalPopulation) * 1000 // Example distribution
	}

	return distribution
}
