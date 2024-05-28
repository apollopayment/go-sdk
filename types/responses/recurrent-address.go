package responses

type RecurrentAddress struct {
	Id                string  `json:"id"`
	AdvancedBalanceId string  `json:"advancedBalanceId"`
	Currency          string  `json:"currency"`
	Network           string  `json:"network"`
	Address           string  `json:"address"`
	Tag               *string `json:"tag"`
	Balance           string  `json:"balance"`
	Alias             *string `json:"alias"`
	Comment           *string `json:"comment"`
}
