package responses

import "time"

type PersonalUser struct {
	Id                string    `json:"id"`
	ClientId          string    `json:"clientId"`
	ClientEmail       string    `json:"clientEmail"`
	ClientName        string    `json:"clientName"`
	DepositWebhookUrl string    `json:"depositWebhookUrl"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type PersonalUserWithAddresses struct {
	Id                string    `json:"id"`
	ClientId          string    `json:"clientId"`
	ClientEmail       string    `json:"clientEmail"`
	ClientName        string    `json:"clientName"`
	DepositWebhookUrl string    `json:"depositWebhookUrl"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`

	Addresses []PersonalAddress `json:"addresses"`
}
