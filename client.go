package baidupay

import (
	"encoding/json"

	"github.com/chrislzg/baidupay/ecommerce"
	"github.com/chrislzg/baidupay/eto"
)

type Client interface {
	// 调起支付
	UnionOrder(*eto.UnionOrderReq) (*eto.UnionOrderRes, error)
	// 支付回调
	ParseUnionOrderNotify(body []byte) (*eto.UnionOrderNotify, error)
	// 回调响应
	NotifyResponse(error) (string, error)
	// 取消核销
	SyncOrderStatus(*eto.SyncOrderStatusReq) error
	// 申请退款
	ApplyOrderRefund(*eto.ApplyOrderRefundReq) (*eto.ApplyOrderRefundResp, error)
	// 退款回调
	ParseRefundNotify(body []byte) (*eto.RefundNotify, error)
	// 请求业务方退款审核
	ParseRefundAudit(body []byte) (*eto.OrderRefundAuditNotify, error)
}

type baiduPayConfig struct {
	DealID     string          `json:"deal_id"`
	AppID      string          `json:"app_id"`
	AppKey     string          `json:"app_key"`
	PrivateKey json.RawMessage `json:"private_key"`
	PublicKey  json.RawMessage `json:"Public_key"`

	WhiteClientIP []string `json:"white_client_ip"`
	DebugMode     bool     `json:"debug_mode"`
	IsDefault     bool     `json:"default"` // 是否是默认客户端
}

func NewClient(confStr json.RawMessage) (Client, error) {
	conf := &baiduPayConfig{}
	err := json.Unmarshal(confStr, conf)
	if err != nil {
		return nil, err
	}
	c := &ecommerce.PayClient{
		DealID:     conf.DealID,
		AppID:      conf.AppID,
		AppKey:     conf.AppKey,
		PrivateKey: conf.PrivateKey,
		PublicKey:  conf.PublicKey,
	}
	return c, nil
}
