package core

// SignFieldsRangeType 用于区分验签字段范围，signFieldsRange 的值：
type SignFieldsRangeType int

const (
	SignFieldsRangeTypeBasic  SignFieldsRangeType = 0 // 0：原验签字段 appKey+dealId+tpOrderId；
	SignFieldsRangeTypeAmount SignFieldsRangeType = 1 //  1：包含 totalAmount 的验签，验签字段包括 appKey+dealId+tpOrderId+totalAmount，固定值为 1
)

type ApiUrl string

const (
	SyncOrderStatusUrl  ApiUrl = "https://etrade-api.baidu.com/cashier/syncOrderStatus"
	ApplyOrderRefundUrl ApiUrl = "https://etrade-api.baidu.com/cashier/applyOrderRefund"
)
