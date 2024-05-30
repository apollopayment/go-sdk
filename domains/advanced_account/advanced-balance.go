package advanced_account

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) AdvancedBalance(ctx context.Context, advancedBalanceId string) responses.BaseResponseGeneric[*responses.AdvancedBalance] {
	var res responses.AdvancedBalance

	_res := d.requester.Request(ctx, "advanced-balance", map[string]string{"advancedBalanceId": advancedBalanceId}, &res)

	return responses.ConvertBase[*responses.AdvancedBalance](_res)
}
