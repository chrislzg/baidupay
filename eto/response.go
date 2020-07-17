package eto

import (
	"encoding/json"
	"net/url"
)

type CallbackModel interface {
	Filled(url.Values)
}

type BaseResponse struct {
	Errno int32           `json:"errno"`
	Msg   string          `json:"msg"`
	Data  json.RawMessage `json:"data"`
}
