package personal_addresses

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) GetUserAddresses(ctx context.Context, payload requests.GetUserAddresses) responses.BaseResponseGeneric[*[]responses.PersonalAddress] {
	res := []responses.PersonalAddress{}

	_res := d.requester.Request(ctx, "personal-addresses/get-user-addresses", payload.ToMap(), &res)

	return responses.ConvertBase[*[]responses.PersonalAddress](_res)
}
