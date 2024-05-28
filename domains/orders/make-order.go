package orders

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) MakeOrder(ctx context.Context, payload requests.CreateOrder) responses.BaseResponseGeneric[*responses.CreateOrder] {
	var res responses.CreateOrder

	_res := d.requester.Request(ctx, "make-order", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.CreateOrder](_res)
}
