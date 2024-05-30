package withdrawals

import (
	"context"
	"github.com/apollopayment/go-sdk/types/requests"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) WithdrawalFeeToken(ctx context.Context, payload requests.WithdrawalFeeToken) responses.BaseResponseGeneric[*responses.WithdrawalFeeToken] {
	var res responses.WithdrawalFeeToken

	_res := d.requester.Request(ctx, "withdrawal-fee-token", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.WithdrawalFeeToken](_res)
}
