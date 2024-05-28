package responses

type PersonalAddress struct {
	Id       string  `json:"id"`
	UserId   string  `json:"userId"`
	Currency string  `json:"currency"`
	Network  string  `json:"network"`
	Address  string  `json:"address"`
	Tag      *string `json:"tag"`
	Balance  string  `json:"balance"`
	IsActive bool    `json:"isActive"`
}
