package withdrawals

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) MakeWithdrawal(ctx context.Context, payload requests.CreateWithdrawal) responses.BaseResponseGeneric[*responses.Withdrawal] {
	var res responses.Withdrawal

	_res := d.requester.Request(ctx, "make-withdrawal", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.Withdrawal](_res)
}
