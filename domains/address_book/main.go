package address_book

import "apollopayment_sdk/requester"

type Domain struct {
	requester *requester.Requester
}

func New(requester *requester.Requester) *Domain {
	return &Domain{
		requester: requester,
	}
}
