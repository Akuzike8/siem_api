package dto


type WazuhHostRestartBody struct {
	Hosts string
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