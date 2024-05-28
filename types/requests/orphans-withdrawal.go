package requests

type OrphansWithdrawal struct {
	Token      string `json:"token"`
	Address    string `json:"address"`
	Comment    string `json:"comment"`
	WebhookUrl string `json:"webhookUrl"`
}

func (o OrphansWithdrawal) ToMap() map[string]any {
	payload := map[string]any{}

	if o.Token != "" {
		payload["token"] = o.Token
	}
	if o.Address != "" {
		payload["address"] = o.Address
	}
	if o.Comment != "" {
		payload["comment"] = o.Comment
	}
	if o.WebhookUrl != "" {
		payload["webhookUrl"] = o.WebhookUrl
	}

	return payload
}
