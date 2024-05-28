package crosschain_bridge

import (
	"context"
	"apollopayment_sdk/types/responses"
)

func (d *Domain) Get(ctx context.Context, crosschainBridgeId string) responses.BaseResponseGeneric[*responses.CrosschainBridge] {
	var res responses.CrosschainBridge

	_res := d.requester.Request(ctx, "bridge/get", map[string]string{"id": crosschainBridgeId}, &res)

	return responses.ConvertBase[*responses.CrosschainBridge](_res)
}
