package withdrawals

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) MakeWithdrawalAsync(ctx context.Context, payload requests.CreateAsyncWithdrawal) responses.BaseResponseGeneric[*responses.AsyncWithdrawal] {
	var res responses.AsyncWithdrawal

	_res := d.requester.Request(ctx, "make-withdrawal-async", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AsyncWithdrawal](_res)
}
