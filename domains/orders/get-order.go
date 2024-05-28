package orders

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetOrder(ctx context.Context, orderId string) responses.BaseResponseGeneric[*responses.Order] {
	var res responses.Order

	_res := d.requester.Request(ctx, "order", map[string]string{"orderId": orderId}, &res)

	return responses.ConvertBase[*responses.Order](_res)
}
