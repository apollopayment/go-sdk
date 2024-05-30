package webhooks

import (
	"context"
	"github.com/apollopayment/go-sdk/types/responses"
)

func (d *Domain) Get(ctx context.Context, webhookId string) responses.BaseResponseGeneric[*responses.Webhook] {
	var res responses.Webhook

	_res := d.requester.Request(ctx, "webhooks/get-verbose", map[string]string{"webhookId": webhookId}, &res)

	return responses.ConvertBase[*responses.Webhook](_res)
}
