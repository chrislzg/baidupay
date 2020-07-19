package eto

import (
	"encoding/json"

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
	TpOrderID       string                   `json:"tpOrderID"`
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

type UnionOrderNotify struct {
	UserID         int64 // 百度用户 ID
	OrderID        int64 // 百度平台订单 ID
	UnitPrice      int32 // 单价	Integer	单位：分
	Count          int32
	TotalMoney     int32           // 总金额	Integer	订单的实际金额，单位：分
	PayMoney       int32           // 扣除各种优惠后用户还需要支付的金额，单位：分
	PromoMoney     int32           // 红包支付金额
	HbMoney        int32           // 红包支付金额
	HbBalanceMoney int32           // 余额支付金额
	GiftCardMoney  int32           // 抵用券金额
	DealID         int64           // 百度收银台的财务结算凭证
	PayTime        int32           // 支付完成时间，时间戳	1463037529
	PromoDetail    json.RawMessage // 订单参与的促销优惠的详细信息
	PayType        int32           // 支付渠道值
	PartnerID      int32           // 支付平台标识值
	Status         int32           // 订单支付状态	Integer	1：未支付；2：已支付；-1：订单取消	2
	TpOrderID      string          // 业务方订单号
	ReturnData     json.RawMessage // 业务方透传数据
	RsaSign        string
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
