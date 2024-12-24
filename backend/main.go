package main

import (
	"encoding/json"
	"holyLibrary-backend/database"
	"log"
	"net/http"
	"os/exec"
)

// Response struct for API responses
type Response struct {
	Message string `json:"message"`
}

// handleTestAPI handles test API endpoint
func handleTestAPI(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Backend is working!"}
	json.NewEncoder(w).Encode(response)
}

// applySystemConfig applies required system configurations at runtime
func applySystemConfig() {
	// Disable Transparent Huge Pages (THP)
	cmdTHP := "echo never | sudo tee /sys/kernel/mm/transparent_hugepage/enabled /sys/kernel/mm/transparent_hugepage/defrag"
	err := exec.Command("bash", "-c", cmdTHP).Run()
	if err != nil {
		log.Fatalf("Failed to disable Transparent Huge Pages (THP): %v", err)
	}

	// Set vm.max_map_count
	cmdMaxMap := "sudo sysctl -w vm.max_map_count=1677720"
	err = exec.Command("bash", "-c", cmdMaxMap).Run()
	if err != nil {
		log.Fatalf("Failed to set vm.max_map_count: %v", err)
	}

	log.Println("System configurations applied successfully!")
}

func main() {
	// Apply system configurations
	applySystemConfig()

	// Initialize databases
	database.ConnectDB()
	database.ConnectMongoDB()
	database.ConnectElasticsearch()

	log.Println("Database connections established")

	// Define routes
	http.HandleFunc("/api/test", handleTestAPI)
	http.HandleFunc("/upload", UploadCBZ)

	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
