package responses

const (
	AddressTypePayIn     string = "PAY_IN"
	AddressTypeBusiness  string = "BUSINESS"
	AddressTypePayOut    string = "PAY_OUT"
	AddressTypePersonal  string = "PERSONAL"
	AddressTypeRecurrent string = "RECURRENT"
)

type ExtendedAddress struct {
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Alias      *string     `json:"alias"`
	Currency   *string     `json:"currency"`
	Network    string      `json:"network"`
	Balance    string      `json:"balance"`
	Address    string      `json:"address"`
	Tag        *string     `json:"tag"`
	Meta       interface{} `json:"meta"`
	IsArchived bool        `json:"isArchived"`
}
