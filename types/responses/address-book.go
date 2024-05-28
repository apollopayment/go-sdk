package responses

type AddressBookAddress struct {
	Id       string   `json:"id"`
	Networks []string `json:"networks"`
	Address  string   `json:"address"`
	Alias    *string  `json:"alias"`
	Comment  *string  `json:"comment"`
}

type AddressBookList struct {
	Addresses []AddressBookAddress `json:"addresses"`
}
