package requests

type CreateAutoSwap struct {
	Address     string `json:"address"`
	Currency    string `json:"currency"`
	Network     string `json:"network"`
	AmountFrom  string `json:"amountFrom"`
	AmountTo    string `json:"amountTo"`
	FeeInAmount bool   `json:"feeInAmount"`
	WebhookUrl  string `json:"webhookUrl"`
}

func (c CreateAutoSwap) ToMap() map[string]any {
	payload := map[string]any{}

	if c.Address != "" {
		payload["address"] = c.Address
	}
	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.Network != "" {
		payload["network"] = c.Network
	}
	if c.AmountFrom != "" {
		payload["amountFrom"] = c.AmountFrom
	}
	if c.AmountTo != "" {
		payload["amountTo"] = c.AmountTo
	}
	if c.WebhookUrl != "" {
		payload["webhookUrl"] = c.WebhookUrl
	}

	payload["feeInAmount"] = c.FeeInAmount

	return payload
}
