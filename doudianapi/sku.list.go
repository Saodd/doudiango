package doudianapi

import (
	"context"
	"errors"
	"github.com/saodd/alog"
)

// SkuList 获取商品sku列表
// https://op.jinritemai.com/docs/api-docs/14/82
func (client *Client) SkuList(ctx context.Context, params SkuListRequest, session string) (*SkuListResponse, error) {
	if err := params.Valid(); err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	var res struct {
		Data *SkuListResponse `json:"data"`
		SystemError
	}
	err := client.Do(ctx, params, "sku.list", session, &res)
	if err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	return res.Data, nil
}

type SkuListRequest map[string]interface{}

func (r SkuListRequest) ToSignMap() map[string]interface{} {
	return r
}
func (r SkuListRequest) Valid() error {
	if v := r["product_id"]; v == "" {
		return errors.New("请传入product_id")
	}
	return nil
}

type SkuListResponse = []struct {
	Id                int64       `json:"id"`
	OpenUserId        int         `json:"open_user_id"`
	OutSkuId          int         `json:"out_sku_id"`
	ProductId         int64       `json:"product_id"`
	ProductIdStr      string      `json:"product_id_str"`
	SpecDetailId1     int64       `json:"spec_detail_id1"`
	SpecDetailId2     int64       `json:"spec_detail_id2"`
	SpecDetailId3     int         `json:"spec_detail_id3"`
	SpecDetailName1   string      `json:"spec_detail_name1"`
	SpecDetailName2   string      `json:"spec_detail_name2"`
	SpecDetailName3   string      `json:"spec_detail_name3"`
	CustomsReportInfo interface{} `json:"customs_report_info"`
	StockNum          int         `json:"stock_num"`
	Price             int         `json:"price"`
	SettlementPrice   int         `json:"settlement_price"`
	SpecId            int64       `json:"spec_id"`
	CreateTime        int         `json:"create_time"`
	Code              string      `json:"code"`
	SkuType           int         `json:"sku_type"`
	StockMap          struct {
	} `json:"stock_map"`
	StepStockNum        int         `json:"step_stock_num"`
	PromStockNum        int         `json:"prom_stock_num"`
	PromStepStockNum    int         `json:"prom_step_stock_num"`
	PreholdStockNum     int         `json:"prehold_stock_num"`
	PreholdStepStockNum int         `json:"prehold_step_stock_num"`
	PreholdStockMap     interface{} `json:"prehold_stock_map"`
	IsSuit              int         `json:"is_suit"`
	SuitNum             int         `json:"suit_num"`
	Volume              int         `json:"volume"`
	NormalStockNum      int         `json:"normal_stock_num"`
	ChannelStockNum     int         `json:"channel_stock_num"`
}
