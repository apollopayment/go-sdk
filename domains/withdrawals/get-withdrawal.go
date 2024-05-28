package withdrawals

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) GetWithdrawal(ctx context.Context, withdrawalId string) responses.BaseResponseGeneric[*responses.Withdrawal] {
	var res responses.Withdrawal

	_res := d.requester.Request(ctx, "get-withdrawal", map[string]string{"withdrawalId": withdrawalId}, &res)

	return responses.ConvertBase[*responses.Withdrawal](_res)
}
