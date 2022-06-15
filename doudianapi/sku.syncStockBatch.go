package doudianapi

import (
	"context"
	"github.com/saodd/alog"
)

// SkuSyncStockBatch 批量同步接口
// https://op.jinritemai.com/docs/api-docs/14/298
func (client *Client) SkuSyncStockBatch(ctx context.Context, params SkuSyncStockBatchRequest, session string) (*SkuSyncStockBatchResponse, error) {
	if err := params.Valid(); err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	var res struct {
		Data *SkuSyncStockBatchResponse `json:"data"`
		SystemError
	}
	err := client.Do(ctx, params, "sku.syncStockBatch", session, &res)
	if err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	return res.Data, nil
}

type SkuSyncStockBatchRequest map[string]interface{}

func (r SkuSyncStockBatchRequest) ToSignMap() map[string]interface{} {
	return r
}
func (r SkuSyncStockBatchRequest) Valid() error {
	return nil
}

type SkuSyncStockBatchResponse struct {
}
