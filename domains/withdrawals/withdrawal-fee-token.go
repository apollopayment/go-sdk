package withdrawals

import (
	"context"
	"apollopayment_sdk/types/requests"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) WithdrawalFeeToken(ctx context.Context, payload requests.WithdrawalFeeToken) responses.BaseResponseGeneric[*responses.WithdrawalFeeToken] {
	var res responses.WithdrawalFeeToken

	_res := d.requester.Request(ctx, "withdrawal-fee-token", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.WithdrawalFeeToken](_res)
}
