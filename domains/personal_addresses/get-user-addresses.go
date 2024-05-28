package personal_addresses

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetUserAddresses(ctx context.Context, payload requests.GetUserAddresses) responses.BaseResponseGeneric[*[]responses.PersonalAddress] {
	res := []responses.PersonalAddress{}

	_res := d.requester.Request(ctx, "personal_addresses/get-user-addresses", payload.ToMap(), &res)

	return responses.ConvertBase[*[]responses.PersonalAddress](_res)
}
