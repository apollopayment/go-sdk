package requests

type CrosschainBridge struct {
	AdvancedBalanceId string `json:"advancedBalanceId"`
	ClientId          string `json:"clientId"`
	AddressFromId     string `json:"addressFromId"`
	AddressToId       string `json:"addressToId"`
	FeeToken          string `json:"feeToken"`
	WebhookUrl        string `json:"webhookUrl"`
}

func (c CrosschainBridge) ToMap() map[string]any {
	payload := map[string]any{}

	if c.AdvancedBalanceId != "" {
		payload["advancedBalanceId"] = c.AdvancedBalanceId
	}
	if c.ClientId != "" {
		payload["clientId"] = c.ClientId
	}
	if c.AddressFromId != "" {
		payload["addressFromId"] = c.AddressFromId
	}
	if c.AddressToId != "" {
		payload["addressToId"] = c.AddressToId
	}
	if c.FeeToken != "" {
		payload["feeToken"] = c.FeeToken
	}
	if c.WebhookUrl != "" {
		payload["webhookUrl"] = c.WebhookUrl
	}

	return payload
}
