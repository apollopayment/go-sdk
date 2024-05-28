package responses

type ValidateAddress struct {
	IsValid bool   `json:"isValid"`
	Regex   string `json:"regex"`
}
