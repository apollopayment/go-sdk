package personal_addresses

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetUserAddress(ctx context.Context, payload requests.GetUserAddress) responses.BaseResponseGeneric[*responses.PersonalAddress] {
	var res responses.PersonalAddress

	_res := d.requester.Request(ctx, "personal_addresses/get-user-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PersonalAddress](_res)
}
