package handles

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Akuzike8/siem_api/config"
	"github.com/Akuzike8/siem_api/dto"
)

func WazuhHostRestart(w http.ResponseWriter, r *http.Request) {
	var body dto.WazuhHostRestartBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	hosts := strings.Split(body.Hosts, ",")
	token := fmt.Sprintf("Bearer %s",body.Token)
	fields := "id,name"
	var hostIds []string
	var host_Ids string

	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}
	
	for _, host := range hosts{
		url := fmt.Sprintf("%s/agents?name=%s&select=%s",wazuh_connection_string,host,fields)
		
		// Build the GET request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v", err)
		}
	
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)
	
		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		var agents dto.WazuhHostRestartAgentRes

		if err := json.NewDecoder(resp.Body).Decode(&agents); err != nil {
			fmt.Printf("Failed to parse JSON: %v", err)
		}

		hostIds = append(hostIds, agents.Data.AffectedItems[0].Id)
	}
	host_Ids = strings.Join(hostIds, ",")
	
	url := fmt.Sprintf("%s/agents/restart?agents_list=%s",wazuh_connection_string,host_Ids)

	// Build the GET request
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var host_ids dto.WazuhHostRestartAgentRess

	if err := json.NewDecoder(resp.Body).Decode(&host_ids); err != nil {
		fmt.Printf("Failed to parse JSON: %v", err)
	}

	res := fmt.Sprintf("restarted: %s", host_ids.Data.AffectedItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"message": res})
	
}