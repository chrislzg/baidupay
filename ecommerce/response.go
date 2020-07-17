package ecommerce

import (
	"encoding/json"
	"net/url"

	"baidupay/core"
	"baidupay/eto"
)

func (c *PayClient) NotifyResponse(err error) (string, error) {
	baseResponse := &eto.BaseResponse{}
	if err != nil {
		baseResponse = &eto.BaseResponse{
			Errno: 1,
			Msg:   err.Error(),
		}
	}

	res, e := json.Marshal(baseResponse)
	if e != nil {
		return "", e
	}
	return string(res), nil
}

// 验证回调是否是合法的，防止第三方伪造
func (c *PayClient) validCallback(values url.Values) bool {
	sign := values.Get("rsaSign")
	values.Del("rsaSign")

	signErr := core.CheckSign(values.Encode(), sign, c.PublicKey)
	if signErr != nil {
		return false
	}
	return true
}

func (c *PayClient) VerifyNotify(body []byte) error {
	var fieldMap map[string]interface{}
	err := json.Unmarshal(body, &fieldMap)
	if err != nil {
		return err
	}
	plainString := core.BuildSignatureString(fieldMap)
	sign := fieldMap["rsaSign"].(string)
	return core.CheckSign(plainString, sign, c.PublicKey)
}

// 解析回调通知内容，res必须为指针类型
func (c *PayClient) ParseNotify(body []byte, res interface{}) error {
	if err := c.VerifyNotify(body); err != nil {
		return err
	}
	baseResponse := &eto.BaseResponse{}
	err := json.Unmarshal(body, baseResponse)
	if err != nil {
		return err
	}
	if baseResponse.Errno != 0 {
		return &core.ResponseErr{
			ErrNo: baseResponse.Errno,
			Msg:   baseResponse.Msg,
		}
	}
	err = json.Unmarshal(baseResponse.Data, res)
	if err != nil {
		return err
	}
	return nil
}
