package ecommerce

import (
    "fmt"
    "testing"
)

func Test_UnionOrderNotify(t *testing.T) {
    c := new(PayClient)
    s := "" // TODO
    res, err := c.ParseUnionOrderNotify([]byte(s))
    fmt.Println(res)
    fmt.Println(err)
}

func Test_ParseRefundNotify(t *testing.T) {
    c := new(PayClient)
    s := "" // TODO
    res, err := c.ParseRefundNotify([]byte(s))
    fmt.Println(res)
    fmt.Println(err)
}

func Test_ParseRefundAudit(t *testing.T) {
    c := new(PayClient)
    s := "" // TODO
    res, err := c.ParseRefundAudit([]byte(s))
    fmt.Println(err)
    fmt.Println(res)
}
