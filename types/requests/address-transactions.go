package requests

type AddressTransactions struct {
	Id     string   `json:"id"`
	Status []string `json:"status"`
	Type   string   `json:"type"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

func (a AddressTransactions) ToMap() map[string]any {
	payload := map[string]any{}

	if a.Id != "" {
		payload["id"] = a.Id
	}
	if a.Type != "" {
		payload["type"] = a.Type
	}
	if a.Limit > 0 {
		payload["limit"] = a.Limit
	} else {
		payload["limit"] = 100
	}
	if a.Offset >= 0 {
		payload["offset"] = a.Offset
	} else {
		payload["offset"] = 0
	}

	if a.Status != nil && len(a.Status) > 0 {
		payload["status"] = a.Status
	}

	return payload
}
