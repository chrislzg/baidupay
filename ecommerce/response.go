package ecommerce

import (
	"encoding/json"
	"net/url"

	"github.com/chrislzg/baidupay/core"
	"github.com/chrislzg/baidupay/eto"
)

func (c *PayClient) NotifyResponse(err error) (string, error) {
	baseResponse := &eto.BaseResponse{
		Errno: 0,
		Msg:   "success",
	}
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

func (c *PayClient) PayNotifyResponse(err error) (string, error) {
	baseResponse := &eto.BaseResponse{
		Errno: 0,
		Msg:   "success",
		Data:  json.RawMessage(`{"isConsumed": 2}`),
	}
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

	signErr := core.CheckSign(values.Encode(), sign, c.PlatformRsaPublicKey)
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
	return core.CheckSign(plainString, sign, c.PlatformRsaPublicKey)
}

func (c *PayClient) parseGetCallBack(res eto.CallbackModel, body string) error {
	qs, err := url.ParseQuery(body)
	if err == nil {
		res.Filled(qs)
	}

	if !c.validCallback(qs) {
		return core.ErrorInvalidSign
	}
	return nil
}

// 解析回调通知内容，res必须为指针类型
func (c *PayClient) ParseNotify(body []byte, res eto.CallbackModel) error {
	err := c.parseGetCallBack(res, string(body))
	if err != nil {
		return err
	}
	return nil
}
