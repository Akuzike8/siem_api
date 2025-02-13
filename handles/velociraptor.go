package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/Akuzike8/siem_api/config"
	"github.com/Akuzike8/siem_api/connections"
	"github.com/Akuzike8/siem_api/dto"
)

func VelociraptorQuarantine(w http.ResponseWriter, r *http.Request){
	var body dto.VelociraptorHostsBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	// Read config file path from environment or command-line arguments
	cfg := config.LoadConfig()

	configPath := cfg.VEL_CONFIG_PATH
	if configPath == "" {
		log.Fatal("Config path is not provided")
	}

	// Load configuration
	config, err := connections.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Establish a gRPC connection
	conn, err := connections.CreateGRPCConnection(config)
	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	defer conn.Close()

	// Execute a VQL query
	query := "SELECT * FROM info()"

	queryClient, err := connections.ExecuteVQLQuery(conn, query)
	if err != nil {
        http.Error(w, fmt.Sprintf("Error executing VQL query: %v", err), http.StatusInternalServerError)
        return
    }

    // Process the responses
    if err := connections.ProcessResponses(queryClient); err != nil {
        http.Error(w, fmt.Sprintf("Error processing responses: %v", err), http.StatusInternalServerError)
        return
    }

	// Return success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Query executed successfully"))

}

func VelociraptorUnQuarantine(w http.ResponseWriter, r *http.Request){
	var body dto.VelociraptorHostsBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	// Read config file path from environment or command-line arguments
	configPath := os.Getenv("VEL_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path is not provided")
	}

	// Load configuration
	config, err := connections.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Establish a gRPC connection
	conn, err := connections.CreateGRPCConnection(config)
	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	defer conn.Close()

	// Execute a VQL query
	query := "SELECT * FROM info()"
	_, err = connections.ExecuteVQLQuery(conn, query)
	if err != nil {
		log.Fatalf("Error executing VQL query: %v", err)
	}
}