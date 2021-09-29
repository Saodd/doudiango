package doudianapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_DoudianProductListV2(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    DoudianProductListV2Request
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
				data: DoudianProductListV2Request{
					"page": 1,
					"size": 2,
				},
				session: secrets.Shop1.AccessToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.DoudianProductListV2(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoudianProductListV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				j, _ := json.Marshal(got)
				fmt.Println(string(j))
			}
		})
	}
}
