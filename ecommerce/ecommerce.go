package ecommerce

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/chrislzg/baidupay/core"
	"github.com/chrislzg/baidupay/eto"
)

type PayClient struct {
	DealID     string
	AppKey     string
	PrivateKey []byte
	PublicKey  []byte
	// ip白名单，回调校验
	WhiteClientIP map[string]bool

	// 调试默认，会输出请求的上下文信息
	debugMode bool
	isDefault bool

	httpClient *http.Client
}

func (c *PayClient) RequestApi() {

}
func (c *PayClient) doRequest(requestData interface{}, url core.ApiUrl, httpMethod string) ([]byte, error) {
	var data []byte
	if requestData != nil {
		var err error
		data, err = json.Marshal(requestData)
		if err != nil {
			return nil, err
		}
	}
	resp, err := core.SimpleRequest(c.httpClient, string(url), httpMethod, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respData := &eto.BaseResponse{}
	err = json.Unmarshal(body, respData)
	if err != nil {
		return nil, err
	}
	if respData.Errno != 0 {
		return nil, &core.ResponseErr{
			ErrNo: respData.Errno,
			Msg:   respData.Msg,
		}
	}
	return body, nil
}
