package doudianapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_SkuSyncStockBatch(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    SkuSyncStockBatchRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动用例",
			args: args{
				ctx: context.Background(),
				data: SkuSyncStockBatchRequest{
					"product_id":  "3495121972449417403",
					"incremental": false,
					"sku_sync_list": map[string]interface{}{
						"sku_id":    1706959605230640,
						"sku_type":  0,
						"stock_num": 6,
					},
				},
				session: secrets.Shop1.AccessToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.SkuSyncStockBatch(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkuSyncStockBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				j, _ := json.Marshal(got)
				fmt.Println(string(j))
			}
		})
	}
}
