package doudianapi

import (
	"context"
	"github.com/saodd/alog"
)

// SkuDetail 获取商品sku详情
// 根据sku id获取商品sku详情
// https://op.jinritemai.com/docs/api-docs/14/566
func (client *Client) SkuDetail(ctx context.Context, params SkuDetailRequest, session string) (*SkuDetailResponse, error) {
	if err := params.Valid(); err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	var res struct {
		Data *SkuDetailResponse `json:"data"`
		SystemError
	}
	err := client.Do(ctx, params, "sku.detail", session, &res)
	if err != nil {
		client.HandleError(ctx, err, alog.V{"params": params})
		return nil, err
	}
	return res.Data, nil
}

type SkuDetailRequest map[string]interface{}

func (r SkuDetailRequest) ToSignMap() map[string]interface{} {
	return r
}
func (r SkuDetailRequest) Valid() error {
	return nil
}

type SkuDetailResponse struct {
	Code              string `json:"code"`
	CreateTime        int    `json:"create_time"`
	CustomsReportInfo struct {
		BarCode           interface{} `json:"bar_code"`
		FirstMeasureQty   int         `json:"first_measure_qty"`
		FirstMeasureUnit  interface{} `json:"first_measure_unit"`
		GModel            interface{} `json:"g_model"`
		HsCode            interface{} `json:"hs_code"`
		ReportBrandName   interface{} `json:"report_brand_name"`
		ReportName        interface{} `json:"report_name"`
		SecondMeasureQty  int         `json:"second_measure_qty"`
		SecondMeasureUnit interface{} `json:"second_measure_unit"`
		Unit              interface{} `json:"unit"`
		Usage             interface{} `json:"usage"`
	} `json:"customs_report_info"`
	Id                  int64 `json:"id"`
	IsSuit              int   `json:"is_suit"`
	OpenUserId          int   `json:"open_user_id"`
	OutSkuId            int   `json:"out_sku_id"`
	PreholdStepStockNum int   `json:"prehold_step_stock_num"`
	PreholdStockMap     struct {
	} `json:"prehold_stock_map"`
	PreholdStockNum  int         `json:"prehold_stock_num"`
	Price            int         `json:"price"`
	ProductId        int64       `json:"product_id"`
	ProductIdStr     string      `json:"product_id_str"`
	PromStepStockNum int         `json:"prom_step_stock_num"`
	PromStockNum     int         `json:"prom_stock_num"`
	SettlementPrice  int         `json:"settlement_price"`
	ShipRuleMap      interface{} `json:"ship_rule_map"`
	SkuType          int         `json:"sku_type"`
	SpecDetailId1    int64       `json:"spec_detail_id1"`
	SpecDetailId2    int64       `json:"spec_detail_id2"`
	SpecDetailId3    int         `json:"spec_detail_id3"`
	SpecDetailName1  string      `json:"spec_detail_name1"`
	SpecDetailName2  string      `json:"spec_detail_name2"`
	SpecDetailName3  string      `json:"spec_detail_name3"`
	SpecId           int64       `json:"spec_id"`
	StepStockNum     int         `json:"step_stock_num"`
	StockMap         struct {
	} `json:"stock_map"`
	StockNum   int    `json:"stock_num"`
	SuitNum    int    `json:"suit_num"`
	SupplierId string `json:"supplier_id"`
	Volume     int    `json:"volume"`
}
