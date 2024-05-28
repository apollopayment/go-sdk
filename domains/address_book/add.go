package address_book

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) Add(ctx context.Context, payload requests.AddAddressBook) responses.BaseResponseGeneric[*responses.AddressBookAddress] {
	var res responses.AddressBookAddress

	_res := d.requester.Request(ctx, "address-book/add", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AddressBookAddress](_res)
}
