package dto

type WazuhHostRestartBody struct {
	Hosts string
	Token string
}

type WazuhGetRequestBody struct {
	Token string
}

type WazuhHostRestartAgentRes struct {
	Data struct {
		AffectedItems []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"affected_items"`
	} `json:"data"`
}

type WazuhHostRestartAgentRess struct {
	Data struct {
		AffectedItems []string `json:"affected_items"`
	} `json:"data"`
}

type WazuhScaAgentRes struct {
	Data struct {
		AffectedItems []struct {
			Policy_Id string `json:"policy_id"`
			Name      string `json:"name"`
			Score     int    `json:"score"`
		} `json:"affected_items"`

		TotalAffectedItems int           `json:"total_affected_items"`
		TotalFailedItems   int           `json:"total_failed_items"`
		FailedItems        []interface{} `json:"failed_items"`
	} `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type WazuhAgentSummaryRes struct {
	Data struct {
		Connection struct {
			Active         int `json:"active"`
			Disconnected   int `json:"disconnected"`
			NeverConnected int `json:"never_connected"`
			Pending        int `json:"pending"`
			Total          int `json:"total"`
		} `json:"connection"`

		Configuration struct {
			Synced    int `json:"synced"`
			Total     int `json:"total"`
			NotSynced int `json:"not_synced"`
		} `json:"configuration"`
	} `json:"data"`
}

type WazuhScaAgentRess struct {
	Hostname  string `json:"hostname"`
	Policy_Id string `json:"policy_id"`
	Name      string `json:"name"`
	Score     int    `json:"score"`
}

type WazuhSoftwarePackageRes struct {
	Data struct {
		// AffectedItem represents a software package from a scan.
		AffectedItems []struct {
			Scan struct {
				ID   int    `json:"id"`
				Time string `json:"time"`
			} `json:"scan"`
			Source       string `json:"source"`
			Size         int64  `json:"size"`
			Multiarch    string `json:"multiarch"`
			Section      string `json:"section"`
			Vendor       string `json:"vendor"`
			Name         string `json:"name"`
			Architecture string `json:"architecture"`
			Format       string `json:"format"`
			Version      string `json:"version"`
			InstallTime  string `json:"install_time"`
			Description  string `json:"description"`
			Priority     string `json:"priority"`
			Location     string `json:"location"`
			AgentID      string `json:"agent_id"`
		} `json:"affected_items"`

		TotalAffectedItems int           `json:"total_affected_items"`
		TotalFailedItems   int           `json:"total_failed_items"`
		FailedItems        []interface{} `json:"failed_items"`
	} `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type WazuhSoftwarePackageRess struct {
	Hostname string `json:"hostname"`
	Vendor   string `json:"vendor"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Format   string `json:"format"`
}

type WazuhWindowsUpdateRes struct {
	Data struct {
		// AffectedItem represents a software package from a scan.
		AffectedItems []struct {
			ScanId   int64  `json:"scan_id"`
			Hotfix   string `json:"hotfix"`
			ScanTime string `json:"scan_time"`
			AgentID  string `json:"agent_id"`
		} `json:"affected_items"`

		TotalAffectedItems int           `json:"total_affected_items"`
		TotalFailedItems   int           `json:"total_failed_items"`
		FailedItems        []interface{} `json:"failed_items"`
	} `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type WazuhWindowsUpdateRess struct {
	Hostname string `json:"hostname"`
	Hotfix   string `json:"hotfix"`
	ScanTime string `json:"scan_time"`
}
