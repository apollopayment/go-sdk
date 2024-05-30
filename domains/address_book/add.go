package address_book

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Add(ctx context.Context, payload requests.AddAddressBook) responses.BaseResponseGeneric[*responses.AddressBookAddress] {
	var res responses.AddressBookAddress

	_res := d.requester.Request(ctx, "address-book/add", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AddressBookAddress](_res)
}
