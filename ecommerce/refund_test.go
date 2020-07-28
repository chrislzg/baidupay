package ecommerce

import (
	"fmt"
	"testing"

	"github.com/chrislzg/baidupay/eto"
)

func Test_SyncOrderStatus(t *testing.T) {
	// TODO
	c := new(PayClient)
	err := c.SyncOrderStatus(&eto.SyncOrderStatusReq{})
	fmt.Println(err)
}

func Test_ApplyOrderRefund(t *testing.T) {
	// TODO
	c := new(PayClient)
	res, err := c.ApplyOrderRefund(&eto.ApplyOrderRefundReq{})
	fmt.Println(res)
	fmt.Println(err)
}
