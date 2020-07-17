package eto

type SyncOrderStatusReq struct {
	OrderId int32  `json:"orderId"`
	UserId  int32  `json:"userId"`
	Type    int32  `json:"type"` // 固定值3
	AppKey  string `json:"app_key"`
	RsaSign string `json:"rsa_sign"`
}

func (m *SyncOrderStatusReq) FieldMap() map[string]interface{} {
	return map[string]interface{}{
		"orderId": m.OrderId,
		"userId":  m.UserId,
		"type":    m.Type,
		"appKey":  m.AppKey,
	}
}

type ApplyOrderRefundReq struct {
	OrderId          int64  `json:"orderId"`
	UserId           int64  `json:"userId"`
	RefundType       int32  `json:"refundType"`
	RefundReason     string `json:"refundReason"`
	TpOrderId        string `json:"tpOrderId"`
	ApplyRefundMoney string `json:"applyRefundMoney"`
	BizRefundBatchId string `json:"bizRefundBatchId"`
	AppKey           string `json:"app_key"`
	RsaSign          string `json:"rsa_sign"`
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

type ApplyOrderRefundResp struct {
	RefundBatchId  string `json:"refund_batch_id"`
	RefundPayMoney string `json:"refund_pay_money"`
}

type OrderRefundAuditNotify struct {
	OrderId          int64  `json:"orderId"`
	UserId           int64  `json:"userId"`
	TpOrderId        string `json:"tpOrderId"`
	RefundBatchId    string `json:"refundBatchId"`
	ApplyRefundMoney int32  `json:"applyRefundMoney"`
	RsaSign          string `json:"rsaSign"`
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

type RefundNotify struct {
	UserId        int64        `json:"userId"`
	OrderId       int64        `json:"orderId"`
	TpOrderId     string       `json:"tpOrderId"`
	RefundBatchId string       `json:"refundBatchId"`
	RefundStatus  RefundStatus `json:"refundStatus"`
	RsaSign       string       `json:"rsaSign"`
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
