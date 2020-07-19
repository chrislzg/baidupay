package ecommerce

import (
	"encoding/json"

	"github.com/chrislzg/baidupay/core"
	"github.com/chrislzg/baidupay/eto"
)

func (c *PayClient) UnionOrder(req *eto.UnionOrderReq) (*eto.UnionOrderRes, error) {
	res := &eto.UnionOrderRes{
		DealID:          c.DealID,
		AppKey:          c.AppKey,
		TotalAmount:     req.TotalAmount,
		DealTitle:       req.DealTitle,
		SignFieldsRange: core.SignFieldsRangeTypeAmount,
		TpOrderID:       req.TpOrderID,
	}

	cipherText, err := core.Sign(res, c.PrivateKey)
	if err != nil {
		return nil, err
	}

	bizInfo := &eto.BizInfo{
		TpData: &eto.BizInfoTpData{
			AppKey:      c.AppKey,
			DealID:      c.DealID,
			RsaSign:     cipherText,
			TpOrderID:   req.TpOrderID,
			TotalAmount: req.TotalAmount,
			ReturnData:  req.ReturnData,
			DealTitle:   res.DealTitle,
		},
	}

	bs, err := json.Marshal(bizInfo)
	if err != nil {
		return nil, err
	}

	res.RsaSign = cipherText
	res.BizInfo = string(bs)
	return res, nil
}

func (c *PayClient) ParseUnionOrderNotify(body []byte) (*eto.UnionOrderNotify, error) {
	res := &eto.UnionOrderNotify{}
	err := c.ParseNotify(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
