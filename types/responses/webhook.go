package responses

import "time"

const (
	WebhookEventOrder              string = "ORDER"
	WebhookEventWithdrawal         string = "WITHDRAWAL"
	WebhookEventDepositPersonal    string = "DEPOSIT_PERSONAL"
	WebhookEventDepositOrphan      string = "DEPOSIT_ORPHAN"
	WebhookEventCrosschainTransfer string = "CROSSCHAIN_TRANSFER"
	WebhookEventCrosschainSwap     string = "CROSSCHAIN_SWAP"
	WebhookEventRecBillingLink     string = "REC_BILLING_LINK"
	WebhookEventRecSubscription    string = "REC_SUBSCRIPTION"
	WebhookEventRecFreePayment     string = "REC_FREE_PAYMENT"
	WebhookEventFiatDeposit        string = "FIAT_DEPOSIT"
)

const (
	WebhookResponseStatusError   string = "ERROR"
	WebhookResponseStatusSuccess string = "SUCCESS"
)

type Webhook struct {
	Id             string            `json:"id"`
	Event          string            `json:"event"`
	EventId        string            `json:"eventId"`
	RequestUrl     string            `json:"requestUrl"`
	RequestHeaders map[string]string `json:"requestHeaders"`
	RequestBody    string            `json:"requestBody"`
	ResponseStatus string            `json:"responseStatus"`
	ResponseBody   string            `json:"responseBody"`
	SentAt         time.Time         `json:"sentAt"`
	Signature      string            `json:"signature"`
	ApiKey         string            `json:"apiKey"`
	ApiKeyAlias    string            `json:"apiKeyAlias"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}
