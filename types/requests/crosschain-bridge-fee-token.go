package requests

type CrosschainBridgeFeeToken struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	Currency          string `json:"currency"`
	NetworkFrom       string `json:"networkFrom"`
	NetworkTo         string `json:"networkTo"`
	Amount            string `json:"amount"`
}

func (c CrosschainBridgeFeeToken) ToMap() map[string]any {
	payload := map[string]any{}

	if c.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = c.AdvancedBalanceId
	}
	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.NetworkFrom != "" {
		payload["networkFrom"] = c.NetworkFrom
	}
	if c.NetworkTo != "" {
		payload["networkTo"] = c.NetworkTo
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}

	return payload
}
