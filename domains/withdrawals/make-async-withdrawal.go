package withdrawals

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) MakeWithdrawalAsync(ctx context.Context, payload requests.CreateAsyncWithdrawal) responses.BaseResponseGeneric[*responses.AsyncWithdrawal] {
	var res responses.AsyncWithdrawal

	_res := d.requester.Request(ctx, "make-withdrawal-async", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AsyncWithdrawal](_res)
}
