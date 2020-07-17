package core

import (
	"errors"
	"fmt"
)

var (
	// ErrorInvalidSign 签名校验失败
	ErrorInvalidSign = errors.New("invalid signature")
)

type ResponseErr struct {
	ErrNo int32
	Msg   string
}

func (e *ResponseErr) Error() string {
	return fmt.Sprintf("ErrOn: %d, ErrMsg: %s", e.ErrNo, e.Msg)
}
