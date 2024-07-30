package main

import (
	"log"
	"net/http"
	"smart-water-management/internal/web"
)

func main() {
	web.LoadTemplates()

	http.HandleFunc("/", web.HomePage)
	//http.HandleFunc("/user", web.UserPage)
	http.HandleFunc("/admin", web.AdminPage)
	http.HandleFunc("/api/flowrates", web.GetAreaFlowRates)
	http.HandleFunc("/api/fairdistribution", web.GetFairDistribution)
	http.HandleFunc("/api/leakdetection", web.GetLeakDetection)

	// Serving static files for JavaScript and CSS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
