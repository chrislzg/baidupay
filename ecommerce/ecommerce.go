package ecommerce

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

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

// baidu form
var postFormContentType = "application/x-www-form-urlencoded"

func (c *PayClient) doRequestPostForm(requestData core.SignatureStruct, iUrl core.ApiUrl) ([]byte, error) {
	var data url.Values
	if requestData != nil {
		data = requestData.FieldForm()
	}
	resp, err := core.SimpleRequest(c.httpClient, string(iUrl), http.MethodPost, map[string]string{"Content-Type": postFormContentType}, []byte(data.Encode()))
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
