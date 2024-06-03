package requests

type CreateOrder struct {
	//AdvancedBalanceId string `json:"advancedBalanceId"`
	Currency       string `json:"currency"`
	Network        string `json:"network"`
	Amount         string `json:"amount"`
	ErrorWebhook   string `json:"errorWebhook"`
	SuccessWebhook string `json:"successWebhook"`
	ReturnUrl      string `json:"returnUrl"`
	Order          string `json:"order"`
	Description    string `json:"description"`
	Lifetime       int    `json:"lifetime"`
	CheckRisks     bool   `json:"checkRisks"`
}

func (c CreateOrder) ToMap() map[string]any {
	payload := map[string]any{}

	//if c.AdvancedBalanceId != "" {
	//	payload["advancedBalanceId"] = c.AdvancedBalanceId
	//}
	if c.Order != "" {
		payload["order"] = c.Order
	}
	if c.Description != "" {
		payload["description"] = c.Description
	}
	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.Network != "" {
		payload["network"] = c.Network
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}
	if c.ErrorWebhook != "" {
		payload["errorWebhook"] = c.ErrorWebhook
	}
	if c.SuccessWebhook != "" {
		payload["successWebhook"] = c.SuccessWebhook
	}
	if c.ReturnUrl != "" {
		payload["returnURL"] = c.ReturnUrl
	}
	if c.Lifetime > 0 {
		payload["lifetime"] = c.Lifetime
	} else {
		switch c.Network {
		case "ripple", "bsc", "tron", "ethereum", "fantom":
			payload["lifetime"] = 1800
			break
		case "litecoin":
			payload["lifetime"] = 3600
			break
		case "bitcoin", "bitcoincash":
			payload["lifetime"] = 7200
			break
		}
	}

	payload["checkRisks"] = c.CheckRisks

	return payload
}
