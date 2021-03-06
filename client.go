package baidupay

import (
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
	// 支付回调响应
	PayNotifyResponse(err error) (string, error)
}

type PayConfig struct {
	DealID               string `json:"deal_id"`
	AppID                string `json:"app_id"`
	AppKey               string `json:"app_key"`
	PrivateKey           []byte `json:"private_key"`
	PublicKey            []byte `json:"public_key"`
	PlatformRsaPublicKey []byte `json:"platform_rsa_public_key"`

	WhiteClientIP []string `json:"white_client_ip"`
	DebugMode     bool     `json:"debug_mode"`
	IsDefault     bool     `json:"default"` // 是否是默认客户端
}

func NewClient(conf *PayConfig) (Client, error) {
	c := &ecommerce.PayClient{
		DealID:               conf.DealID,
		AppKey:               conf.AppKey,
		PrivateKey:           conf.PrivateKey,
		PlatformRsaPublicKey: conf.PlatformRsaPublicKey,
		PublicKey:            conf.PublicKey,
	}
	return c, nil
}
