package requests

type CreatePersonalUser struct {
	ClientId          string `json:"clientId"`
	ClientEmail       string `json:"clientEmail"`
	ClientName        string `json:"clientName"`
	DepositWebhookUrl string `json:"depositWebhookUrl"`
	CreateAddresses   bool   `json:"createAddresses"`
	GroupByBlockchain bool   `json:"groupByBlockchain"`
	CheckRisks        bool   `json:"checkRisks"`
}

func (c CreatePersonalUser) ToMap() map[string]any {
	payload := map[string]any{}

	if c.ClientId != "" {
		payload["clientId"] = c.ClientId
	}
	if c.ClientEmail != "" {
		payload["clientEmail"] = c.ClientEmail
	}
	if c.ClientName != "" {
		payload["clientName"] = c.ClientName
	}
	if c.DepositWebhookUrl != "" {
		payload["depositWebhookUrl"] = c.DepositWebhookUrl
	}

	payload["createAddresses"] = c.CreateAddresses
	payload["groupByBlockchain"] = c.GroupByBlockchain
	payload["checkRisks"] = c.CheckRisks

	return payload
}
