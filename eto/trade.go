package eto

import (
	"encoding/json"
    "net/url"

    "github.com/chrislzg/baidupay/core"
)

// https://dianshang.baidu.com/platform/doclist/index.html#!/doc/nuomiplus_1_guide/mini_program_cashier/parameter.md
type UnionOrderReq struct {
	TotalAmount string // 订单总金额，以分为单位
	TpOrderID   string
	DealTitle   string
	ReturnData  string
}

/*
{
		"dealId": "470193086",
		"appKey": "MMMabc",
		"totalAmount": "11300",
		"tpOrderId": "3028903626",
		"dealTitle": "爱鲜蜂",
		“signFieldsRange”：“1”,
		"rsaSign": '',
		"bizInfo": '{"tpData":{"appKey":"MMMabc","dealId":"470193086","tpOrderId":"3028903626",
			"rsaSign":"","totalAmount":"11300","payResultUrl":"",
			"returnData":{"bizKey1":"第三方的字段1取值","bizKey2":"第三方的字段2取值"}}}'
}
*/
type UnionOrderRes struct {
	DealID          string                   `json:"dealId"`
	AppKey          string                   `json:"appKey"`
	TotalAmount     string                   `json:"totalAmount"` // 订单总金额，以分为单位
	TpOrderID       string                   `json:"tpOrderId"`
	SignFieldsRange core.SignFieldsRangeType `json:"signFieldsRange"`
	DealTitle       string                   `json:"dealTitle"` // 订单的名称
	RsaSign         string                   `json:"rsaSign"`
	BizInfo         string                   `json:"bizInfo"`
}

func (m *UnionOrderRes) FieldMap() map[string]interface{} {
	return map[string]interface{}{
		"dealId":      m.DealID,
		"appKey":      m.AppKey,
		"tpOrderId":   m.TpOrderID,
		"totalAmount": m.TotalAmount,
	}
}

type BizInfo struct {
	TpData *BizInfoTpData `json:"tpData"`
}

type BizInfoTpData struct {
	AppKey         string `json:"appKey"`
	DealID         string `json:"dealID"`
	TpOrderID      string `json:"tpOrderID"`
	RsaSign        string `json:"rsaSign"`
	TotalAmount    string `json:"totalAmount"` // 订单总金额，以分为单位
	PayResultUrl   string `json:"payResultUrl"`
	ReturnData     string `json:"returnData"` // 业务方用于透传的业务变量
	DealTitle      string `json:"dealTitle"`  // 订单的名称
	DetailSubTitle string `json:"detailSubTitle"`
	DealTumbView   string `json:"dealTumbView"`
}

/*
type UnionOrderNotify struct {
    UserID         int64           `json:"userId"`    // 百度用户 ID
    OrderID        int64           `json:"orderId"`   // 百度平台订单 ID
    UnitPrice      int32           `json:"unitPrice"` // 单价	Integer	单位：分
    Count          int32           `json:"count"`
    TotalMoney     int32           `json:"totalMoney"`     // 总金额	Integer	订单的实际金额，单位：分
    PayMoney       int32           `json:"payMoney"`       // 扣除各种优惠后用户还需要支付的金额，单位：分
    PromoMoney     int32           `json:"promoMoney"`     // 红包支付金额
    HbMoney        int32           `json:"hbMoney"`        // 红包支付金额
    HbBalanceMoney int32           `json:"hbBalanceMoney"` // 余额支付金额
    GiftCardMoney  int32           `json:"giftCardMoney"`  // 抵用券金额
    DealID         int64           `json:"dealId"`         // 百度收银台的财务结算凭证
    PayTime        int32           `json:"payTime"`        // 支付完成时间，时间戳	1463037529
    PromoDetail    json.RawMessage `json:"promoDetail"`    // 订单参与的促销优惠的详细信息
    PayType        int32           `json:"payType"`        // 支付渠道值
    PartnerID      int32           `json:"partnerId"`      // 支付平台标识值
    Status         string           `json:"status"`         // 订单支付状态	Integer	1：未支付；2：已支付；-1：订单取消	2
    TpOrderID      string          `json:"tpOrderId"`      // 业务方订单号
    ReturnData     json.RawMessage `json:"returnData"`     // 业务方透传数据
    RsaSign        string          `json:"rsaSign"`
}
*/
type UnionOrderNotify struct {
    UserID         string           `json:"userId"`    // 百度用户 ID
    OrderID        string           `json:"orderId"`   // 百度平台订单 ID
    UnitPrice      string           `json:"unitPrice"` // 单价	Integer	单位：分
    Count          string           `json:"count"`
    TotalMoney     string           `json:"totalMoney"`     // 总金额	Integer	订单的实际金额，单位：分
    PayMoney       string           `json:"payMoney"`       // 扣除各种优惠后用户还需要支付的金额，单位：分
    PromoMoney     string           `json:"promoMoney"`     // 红包支付金额
    HbMoney        string           `json:"hbMoney"`        // 红包支付金额
    HbBalanceMoney string           `json:"hbBalanceMoney"` // 余额支付金额
    GiftCardMoney  string           `json:"giftCardMoney"`  // 抵用券金额
    DealID         string           `json:"dealId"`         // 百度收银台的财务结算凭证
    PayTime        string           `json:"payTime"`        // 支付完成时间，时间戳	1463037529
    PromoDetail    json.RawMessage `json:"promoDetail"`    // 订单参与的促销优惠的详细信息
    PayType        string           `json:"payType"`        // 支付渠道值
    PartnerID      string           `json:"partnerId"`      // 支付平台标识值
    Status         string           `json:"status"`         // 订单支付状态	Integer	1：未支付；2：已支付；-1：订单取消	2
    TpOrderID      string          `json:"tpOrderId"`      // 业务方订单号
    ReturnData     json.RawMessage `json:"returnData"`     // 业务方透传数据
    RsaSign        string          `json:"rsaSign"`
}

// 是否支付成功
func (m *UnionOrderNotify) IsSuccess() bool {
    return m.Status == "2"
}

// 通过 values 填充自己
func (m *UnionOrderNotify) Filled(values url.Values) {
    m.UserID = values.Get("userId")
    m.OrderID = values.Get("orderId")
    m.UnitPrice = values.Get("unitPrice")
    m.Count = values.Get("count")
    m.TotalMoney = values.Get("totalMoney")
    m.PayMoney = values.Get("payMoney")
    m.PromoMoney = values.Get("promoMoney")
    m.HbMoney = values.Get("hbMoney")
    m.HbBalanceMoney = values.Get("hbBalanceMoney")
    m.GiftCardMoney = values.Get("giftCardMoney")
    m.DealID = values.Get("dealId")
    m.PayTime = values.Get("payTime")
    m.PromoDetail = json.RawMessage(values.Get("promoDetail"))
    m.PayType = values.Get("payType")
    m.PartnerID = values.Get("partnerId")
    m.Status = values.Get("status")
    m.TpOrderID = values.Get("tpOrderId")
    m.ReturnData = json.RawMessage(values.Get("returnData"))
    m.RsaSign = values.Get("rsaSign")
}

func (m *UnionOrderNotify) FiledMap() map[string]interface{} {
	return map[string]interface{}{
		"userId":         m.UserID,
		"orderId":        m.OrderID,
		"unitPrice":      m.UnitPrice,
		"count":          m.Count,
		"totalMoney":     m.TotalMoney,
		"payMoney":       m.PayMoney,
		"promoMoney":     m.PromoMoney,
		"hbMoney":        m.HbMoney,
		"hbBalanceMoney": m.HbBalanceMoney,
		"giftCardMoney":  m.GiftCardMoney,
		"dealId":         m.DealID,
		"payTime":        m.PayTime,
		"promoDetail":    m.PromoDetail,
		"payType":        m.PayType,
		"partnerId":      m.PartnerID,
		"status":         m.Status,
		"tpOrderId":      m.TpOrderID,
		"returnData":     m.ReturnData,
	}
}
