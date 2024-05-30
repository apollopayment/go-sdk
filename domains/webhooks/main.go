package webhooks

import "github.com/apollopayment/go-sdk/requester"

type Domain struct {
	requester *requester.Requester
}

func New(requester *requester.Requester) *Domain {
	return &Domain{
		requester: requester,
	}
}
