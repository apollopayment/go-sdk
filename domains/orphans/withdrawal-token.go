package orphans

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) WithdrawalToken(ctx context.Context, orphanId string) responses.BaseResponseGeneric[*responses.OrphansWithdrawalToken] {
	var res responses.OrphansWithdrawalToken

	_res := d.requester.Request(ctx, "orphan-deposits/withdrawal-token", map[string]string{"id": orphanId}, &res)

	return responses.ConvertBase[*responses.OrphansWithdrawalToken](_res)
}
