package ecommerce

import (
	"encoding/json"

	"baidupay/core"
	"baidupay/eto"
)

func (c *PayClient) UnionOrder(req *eto.UnionOrderReq) (*eto.UnionOrderRes, error) {
	res := &eto.UnionOrderRes{
		DealID:          c.DealID,
		AppKey:          c.AppKey,
		TotalAmount:     req.Amount,
		DealTitle:       req.ProductName,
		SignFieldsRange: core.SignFieldsRangeTypeAmount,
		TpOrderID:       req.OrderID,
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
			TpOrderID:   req.OrderID,
			TotalAmount: req.Amount,
			ReturnData:  req.Attach,
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
