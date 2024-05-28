package requests

type GetOrders struct {
	Status []string `json:"status"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

func (g GetOrders) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Limit > 0 {
		payload["limit"] = g.Limit
	}
	if g.Offset >= 0 {
		payload["offset"] = g.Offset
	}

	if g.Status != nil && len(g.Status) != 0 {
		payload["status"] = g.Status
	}

	return payload
}
