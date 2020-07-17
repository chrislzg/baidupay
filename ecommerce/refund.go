package ecommerce

import (
	"encoding/json"
	"net/http"

	"baidupay/core"
	"baidupay/eto"
)

func (c *PayClient) SyncOrderStatus(req *eto.SyncOrderStatusReq) error {
	err := c.EnrichSyncOrderStatusReq(req)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, core.SyncOrderStatusUrl, http.MethodPost)
	if err != nil {
		return err
	}
	return nil
}

func (c *PayClient) EnrichSyncOrderStatusReq(req *eto.SyncOrderStatusReq) error {
	req.AppKey = c.AppKey
	req.Type = 3
	rsaSign, err := core.Sign(req, c.PrivateKey)
	if err != nil {
		return err
	}
	req.RsaSign = rsaSign
	return nil
}

func (c *PayClient) ApplyOrderRefund(req *eto.ApplyOrderRefundReq) (*eto.ApplyOrderRefundResp, error) {
	err := c.EnrichApplyOrderRefundReq(req)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.doRequest(req, core.ApplyOrderRefundUrl, http.MethodPost)
	if err != nil {
		return nil, err
	}
	res := &eto.ApplyOrderRefundResp{}
	err = json.Unmarshal(httpResponse, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *PayClient) EnrichApplyOrderRefundReq(req *eto.ApplyOrderRefundReq) error {
	req.AppKey = c.AppKey

	rsaSign, err := core.Sign(req, c.PrivateKey)
	if err != nil {
		return err
	}
	req.RsaSign = rsaSign
	return nil
}

func (c *PayClient) ParseRefundNotify(body []byte) (*eto.RefundNotify, error) {
	res := &eto.RefundNotify{}
	err := c.ParseNotify(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
