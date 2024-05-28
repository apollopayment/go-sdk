package base

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) PriceRate(ctx context.Context, payload requests.PriceRate) responses.BaseResponseGeneric[*string] {
	var res string

	_res := d.requester.Request(ctx, "price-rate", payload.ToMap(), &res)

	return responses.ConvertBase[*string](_res)
}
