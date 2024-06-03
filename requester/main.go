package requester

import "github.com/apollopayment/go-sdk/noncer"

type Requester struct {
	publicKey string
	secretKey string

	noncer *noncer.Noncer

	advancedBalanceId string
}

func New(public, secret string) *Requester {
	return &Requester{
		publicKey: public,
		secretKey: secret,

		noncer: noncer.New(),
	}
}

func (r *Requester) SetAdvancedBalance(adv string) {
	r.advancedBalanceId = adv
}
