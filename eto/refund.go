package eto

import (
	"fmt"
	"net/url"
)

type SyncOrderStatusReq struct {
	OrderId int64  `json:"orderId"`
	UserId  int64  `json:"userId"`
	Type    int32  `json:"type"` // 固定值3
	AppKey  string `json:"appKey"`
	RsaSign string `json:"rsaSign"`
}

func (m *SyncOrderStatusReq) FieldMap() map[string]interface{} {
	return map[string]interface{}{
		"orderId": m.OrderId,
		"userId":  m.UserId,
		"type":    m.Type,
		"appKey":  m.AppKey,
	}
}

func (m *SyncOrderStatusReq) FieldForm() url.Values {
	values := url.Values{}
	fieldMap := m.FieldMap()
	for key := range fieldMap {
		values.Add(key, fmt.Sprint(fieldMap[key]))
	}
	values.Add("rsaSign", m.RsaSign)
	return values
}

type ApplyOrderRefundReq struct {
	OrderId          int64  `json:"orderId"`
	UserId           int64  `json:"userId"`
	RefundType       int32  `json:"refundType"`
	RefundReason     string `json:"refundReason"`
	TpOrderId        string `json:"tpOrderId"`
	ApplyRefundMoney int64  `json:"applyRefundMoney"`
	BizRefundBatchId string `json:"bizRefundBatchId"`
	AppKey           string `json:"appKey"`
	RsaSign          string `json:"rsaSign"`
}

func (m *ApplyOrderRefundReq) FieldMap() map[string]interface{} {
	return map[string]interface{}{
		"orderId":          m.OrderId,
		"userId":           m.UserId,
		"refundType":       m.RefundType,
		"refundReason":     m.RefundReason,
		"tpOrderId":        m.TpOrderId,
		"applyRefundMoney": m.ApplyRefundMoney,
		"bizRefundBatchId": m.BizRefundBatchId,
		"appKey":           m.AppKey,
	}
}

func (m *ApplyOrderRefundReq) FieldForm() url.Values {
	values := url.Values{}
	fieldMap := m.FieldMap()
	for key := range fieldMap {
		values.Add(key, fmt.Sprint(fieldMap[key]))
	}
	values.Add("rsaSign", m.RsaSign)
	return values
}

type ApplyOrderRefundResp struct {
	RefundBatchId  string `json:"refundBatchId"`
	RefundPayMoney string `json:"refundPayMoney"`
}

/*
type OrderRefundAuditNotify struct {
	OrderId          int64  `json:"orderId"`
	UserId           int64  `json:"userId"`
	TpOrderId        string `json:"tpOrderId"`
	RefundBatchId    string `json:"refundBatchId"`
	ApplyRefundMoney int32  `json:"applyRefundMoney"`
	RsaSign          string `json:"rsaSign"`
}
*/
type OrderRefundAuditNotify struct {
	OrderId          string `json:"orderId"`
	UserId           string `json:"userId"`
	TpOrderId        string `json:"tpOrderId"`
	RefundBatchId    string `json:"refundBatchId"`
	ApplyRefundMoney string `json:"applyRefundMoney"`
	RsaSign          string `json:"rsaSign"`
}

// 通过 values 填充自己
func (res *OrderRefundAuditNotify) Filled(values url.Values) {
	res.OrderId = values.Get("orderId")
	res.UserId = values.Get("userId")
	res.TpOrderId = values.Get("tpOrderId")
	res.RefundBatchId = values.Get("refundBatchId")
	res.ApplyRefundMoney = values.Get("applyRefundMoney")
	res.RsaSign = values.Get("rsaSign")
}

type AuditStatus int32

const (
	AuditStatusInvalid AuditStatus = iota
	AuditStatusConfirm
	AuditStatusReject
	AuditStatusUnknow
)

type CalculateRes struct {
	RefundPayMoney int32 `json:"refund_pay_money"`
}

type OrderRefundAuditRes struct {
	AuditStatus  AuditStatus  `json:"auditStatus"`
	CalculateRes CalculateRes `json:"calculate_res"`
}

type RefundStatus int32

const (
	RefundInvalid RefundStatus = iota
	RefundSuccess
	RefundFail
)

/*
type RefundNotify struct {
	UserId        int64        `json:"userId"`
	OrderId       int64        `json:"orderId"`
	TpOrderId     string       `json:"tpOrderId"`
	RefundBatchId string       `json:"refundBatchId"`
	RefundStatus  RefundStatus `json:"refundStatus"`
	RsaSign       string       `json:"rsaSign"`
}
*/
type RefundNotify struct {
	UserId        string `json:"userId"`
	OrderId       string `json:"orderId"`
	TpOrderId     string `json:"tpOrderId"`
	RefundBatchId string `json:"refundBatchId"`
	RefundStatus  string `json:"refundStatus"`
	RsaSign       string `json:"rsaSign"`
}

// 通过 values 填充自己
func (res *RefundNotify) Filled(values url.Values) {
	res.UserId = values.Get("userId")
	res.OrderId = values.Get("orderId")
	res.TpOrderId = values.Get("tpOrderId")
	res.RefundBatchId = values.Get("refundBatchId")
	res.RefundStatus = values.Get("refundStatus")
	res.RsaSign = values.Get("rsaSign")
}

//
//func (m *RefundNotify) FieldMap() map[string]interface{} {
//	return map[string]interface{}{
//		"userId":        m.UserId,
//		"orderId":       m.OrderId,
//		"tpOrderId":     m.TpOrderId,
//		"refundBatchId": m.RefundBatchId,
//		"refundStatus":  m.RefundStatus,
//	}
//}
