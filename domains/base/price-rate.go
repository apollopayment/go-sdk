package base

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) PriceRate(ctx context.Context, from, to string) responses.BaseResponseGeneric[*string] {
	var res string

	_res := d.requester.Request(ctx, "price-rate", map[string]string{"from": from, "to": to}, &res)

	return responses.ConvertBase[*string](_res)
}
