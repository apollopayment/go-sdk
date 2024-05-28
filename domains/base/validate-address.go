package base

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) ValidateAddress(ctx context.Context, payload requests.ValidateAddress) responses.BaseResponseGeneric[*responses.ValidateAddress] {
	var res responses.ValidateAddress

	_res := d.requester.Request(ctx, "utils/validate-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.ValidateAddress](_res)
}
