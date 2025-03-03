package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type HealthResponse struct {
	Services map[string]bool `json:"services"`
}

func checkServiceStatus(service string) bool {
	cmd := exec.Command("systemctl", "is-active", service)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "active"
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	servicesParam := r.URL.Query().Get("services")
	if servicesParam == "" {
		http.Error(w, "Missing 'services' query parameter", http.StatusBadRequest)
		return
	}

	services := strings.Split(servicesParam, ",")
	response := HealthResponse{Services: make(map[string]bool)}
	allHealthy := true

	for _, service := range services {
		status := checkServiceStatus(service)
		response.Services[service] = status
		if !status {
			allHealthy = false
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if allHealthy {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	http.HandleFunc("/health", healthCheckHandler)
	log.Printf("Starting server on port %s...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
