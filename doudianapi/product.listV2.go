package doudianapi

import (
	"context"
	"github.com/saodd/alog"
)

// DoudianProductListV2 获取商品列表新版
// 获取商品列表信息，类似商家后台商品管理列表
// https://op.jinritemai.com/docs/api-docs/14/633
func (client *Client) DoudianProductListV2(ctx context.Context, params DoudianProductListV2Request, session string) (*DoudianProductListV2Response, error) {
	if err := params.Valid(); err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	var res struct {
		Data *DoudianProductListV2Response `json:"data"`
		SystemError
	}
	err := client.Do(ctx, params, "product.listV2", session, &res)
	if err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	return res.Data, nil
}

type DoudianProductListV2Request map[string]interface{}

func (r DoudianProductListV2Request) ToSignMap() map[string]interface{} {
	return r
}
func (r DoudianProductListV2Request) Valid() error {
	return nil
}

type DoudianProductListV2Response struct {
	Data []struct {
		CategoryDetail struct {
			FirstCid    int    `json:"first_cid"`
			FirstCname  string `json:"first_cname"`
			FourthCid   int    `json:"fourth_cid"`
			FourthCname string `json:"fourth_cname"`
			SecondCid   int    `json:"second_cid"`
			SecondCname string `json:"second_cname"`
			ThirdCid    int    `json:"third_cid"`
			ThirdCname  string `json:"third_cname"`
		} `json:"category_detail"`
		CheckStatus      int    `json:"check_status"`
		CosRatio         int    `json:"cos_ratio"`
		CreateTime       int    `json:"create_time"`
		Description      string `json:"description"`
		DiscountPrice    int    `json:"discount_price"`
		Extra            string `json:"extra"`
		Img              string `json:"img"`
		IsPackageProduct bool   `json:"is_package_product"`
		MarketPrice      int    `json:"market_price"`
		Mobile           string `json:"mobile"`
		Name             string `json:"name"`
		OutProductId     int    `json:"out_product_id"`
		OuterProductId   string `json:"outer_product_id"`
		PayType          int    `json:"pay_type"`
		ProductId        int64  `json:"product_id"`
		ProductType      int    `json:"product_type"`
		RecommendRemark  string `json:"recommend_remark"`
		SpecId           int64  `json:"spec_id"`
		Status           int    `json:"status"`
		UpdateTime       int    `json:"update_time"`
	} `json:"data"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
