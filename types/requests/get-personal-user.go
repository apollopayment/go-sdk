package requests

type GetPersonalUser struct {
	Id       string `json:"id"`
	ClientId string `json:"clientId"`
}

func (g GetPersonalUser) ToMap() map[string]any {
	payload := map[string]any{}

	if g.Id != "" {
		payload["id"] = g.Id
	}
	if g.ClientId != "" {
		payload["clientId"] = g.ClientId
	}

	return payload
}
