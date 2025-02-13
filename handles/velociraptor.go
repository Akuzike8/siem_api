package handles

import (
	"encoding/json"
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

	// Execute a VQL query
	query := "SELECT * FROM info()"

	results := connections.ExecuteVQLQuery(query)

	// Return success response
    w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type","Application/json")
    w.Write([]byte(results))

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

}