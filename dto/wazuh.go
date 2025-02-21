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
