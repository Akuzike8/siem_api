package handles

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Akuzike8/siem_api/config"
	"github.com/Akuzike8/siem_api/dto"
)

func WazuhAgentSummary(w http.ResponseWriter, r *http.Request) {
	var body dto.WazuhGetRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("Bearer %s",body.Token)
	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}
	url := fmt.Sprintf("%s/agents/summary/status",wazuh_connection_string)

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

	var summary dto.WazuhAgentSummaryRes

	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		fmt.Printf("Failed to parse JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(summary)
}

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

func WazuhAgentList(w http.ResponseWriter, r *http.Request){
	var body dto.WazuhGetRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("Bearer %s",body.Token)
	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}

	url := fmt.Sprintf("%s/agents?limit=1000",wazuh_connection_string)
	
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(agents.Data.AffectedItems)
}

func WazuhCisPosture(w http.ResponseWriter, r *http.Request){
	var body dto.WazuhGetRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("Bearer %s",body.Token)
	fields := "name,policy_id,score"

	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}
	
	url := fmt.Sprintf("%s/agents?limit=1000",wazuh_connection_string)
	
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

	var scores []dto.WazuhScaAgentRess

	for _, agent := range agents.Data.AffectedItems {
		url = fmt.Sprintf("%s/sca/%s?select=%s",wazuh_connection_string,agent.Id,fields)
		
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
		
		// Check for non-2xx status codes
		
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			fmt.Printf("Received non-2xx status for agent %s: %d", agent.Id, resp.StatusCode)
			continue
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body) // Read all the body content
		if err != nil {
			fmt.Print(err)
		}

		var cis_res dto.WazuhScaAgentRes

		// Parse the JSON into the response struct
		if err := json.Unmarshal([]byte(string(body)), &cis_res); err != nil {
			fmt.Print(err)
		}

		if  cis_res.Data.TotalAffectedItems >= 1 {

			score := dto.WazuhScaAgentRess{
				Hostname: agent.Name,
				Policy_Id: cis_res.Data.AffectedItems[0].Policy_Id,
				Name: cis_res.Data.AffectedItems[0].Name,
				Score: cis_res.Data.AffectedItems[0].Score,
			}
	
			scores = append(scores, score)
		}

	
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(scores)
	

	
}

func WazuhWindowsUpdate(w http.ResponseWriter, r *http.Request){
	var body dto.WazuhGetRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("Bearer %s",body.Token)
	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}
	
	url := fmt.Sprintf("%s/agents?limit=10000",wazuh_connection_string)

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

	var osupdates []dto.WazuhWindowsUpdateRess

	for _, agent := range agents.Data.AffectedItems {
		url = fmt.Sprintf("%s/syscollector/%s/hotfixes",wazuh_connection_string,agent.Id)

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

		var windowsupdates dto.WazuhWindowsUpdateRes

		if err := json.NewDecoder(resp.Body).Decode(&windowsupdates); err != nil {
			fmt.Printf("Failed to parse JSON: %v", err)
		}

		if  windowsupdates.Data.TotalAffectedItems >= 1 {

			for _, soft := range windowsupdates.Data.AffectedItems {
				if soft.Hotfix == "" {
					continue
				}

				pkg := dto.WazuhWindowsUpdateRess{
					Hostname: agent.Name,
					Hotfix: soft.Hotfix,
					ScanTime: soft.ScanTime,
				}

				osupdates = append(osupdates, pkg)
			}

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(osupdates)
}

func WazuhSoftwarePackage(w http.ResponseWriter, r *http.Request){
	var body dto.WazuhGetRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w,"can't decode",http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("Bearer %s",body.Token)
	fields := "name,vendor,version,format"

	cfg := config.LoadConfig()

	wazuh_connection_string := cfg.WAZUH_CONNECTION_STRING

	//skipping tls verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create the HTTP client
	client := &http.Client{Transport: transport}
	
	url := fmt.Sprintf("%s/agents?limit=10000",wazuh_connection_string)

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

	var packages []dto.WazuhSoftwarePackageRess

	for _, agent := range agents.Data.AffectedItems {
		url = fmt.Sprintf("%s/syscollector/%s/packages?limit=1000&select=%s",wazuh_connection_string,agent.Id,fields)

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

		var softwares dto.WazuhSoftwarePackageRes

		if err := json.NewDecoder(resp.Body).Decode(&softwares); err != nil {
			fmt.Printf("Failed to parse JSON: %v", err)
		}

		if  softwares.Data.TotalAffectedItems >= 1 {

			for _, soft := range softwares.Data.AffectedItems {
				if soft.Name == "" {
					continue
				}

				pkg := dto.WazuhSoftwarePackageRess{
					Hostname: agent.Name,
					Name: soft.Name,
					Version: soft.Version,
					Vendor: soft.Vendor,
					Format: soft.Format,
				}

				packages = append(packages, pkg)
			}

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(packages)
}