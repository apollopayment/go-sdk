package requests

type GetOrphansDeposits struct {
	Id      string `json:"id"`
	OrderId string `json:"orderId"`
	Stage   string `json:"stage"`
	Status  string `json:"status"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

func (g GetOrphansDeposits) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Id != "" {
		payload["id"] = g.Id
	}
	if g.OrderId != "" {
		payload["orderId"] = g.OrderId
	}
	if g.Stage != "" {
		payload["stage"] = g.Stage
	}
	if g.Status != "" {
		payload["status"] = g.Status
	}

	if g.Limit > 0 {
		payload["limit"] = g.Limit
	}
	if g.Offset >= 0 {
		payload["offset"] = g.Offset
	}

	return payload
}
