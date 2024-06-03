package requests

type CreateInvoice struct {
	//AdvancedBalanceId string     `json:"advancedBalanceId"`
	ExternalId       string     `json:"externalId"`
	Order            string     `json:"order"`
	Description      string     `json:"description"`
	Currency         string     `json:"currency"`
	Amount           string     `json:"amount"`
	IncludeFee       bool       `json:"includeFee"`
	AdditionalFees   []string   `json:"additionalFees"`
	InsurancePercent string     `json:"insurancePercent"`
	SlippagePercent  string     `json:"slippagePercent"`
	WebhookURL       string     `json:"webhookURL"`
	ReturnURL        string     `json:"returnURL"`
	Lifetime         int        `json:"lifetime"`
	Currencies       [][]string `json:"currencies"`
}

func (c CreateInvoice) ToMap() map[string]any {
	payload := map[string]any{}

	//if c.AdvancedBalanceId != "" {
	//	payload["advancedBalanceId"] = c.AdvancedBalanceId
	//}
	if c.ExternalId != "" {
		payload["externalId"] = c.ExternalId
	}
	if c.Order != "" {
		payload["order"] = c.Order
	}
	if c.Description != "" {
		payload["description"] = c.Description
	}
	if c.Currency != "" {
		payload["currency"] = c.Currency
	}
	if c.Amount != "" {
		payload["amount"] = c.Amount
	}
	if c.InsurancePercent != "" {
		payload["insurancePercent"] = c.InsurancePercent
	}
	if c.SlippagePercent != "" {
		payload["slippagePercent"] = c.SlippagePercent
	}
	if c.WebhookURL != "" {
		payload["webhookURL"] = c.WebhookURL
	}
	if c.ReturnURL != "" {
		payload["returnURL"] = c.ReturnURL
	}
	if c.Lifetime > 0 {
		payload["lifetime"] = c.Lifetime
	}

	payload["includeFee"] = c.IncludeFee

	if c.AdditionalFees != nil && len(c.AdditionalFees) != 0 {
		payload["additionalFees"] = c.AdditionalFees
	}

	if c.Currencies == nil {
		payload["currencies"] = [][]string{}
	}

	return payload
}
