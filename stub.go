package main

import (
	"encoding/json"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"golang.org/x/net/context"
)

func init() {
	comm.RegisterReqNew("Echo", NewEchoRequest)
}

type EchoRequest struct {
	A int    `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

func NewEchoRequest(a []byte) interface{} {
	req := &EchoRequest{}
	err := json.Unmarshal(a, req)
	if err != nil {
		panic(err)
	}

	return req
}

type EchoResponse struct {
	C int    `json:"c,omitempty"`
	D string `json:"d,omitempty"`
}

type Test interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	SendCash()
	GetOrder()
}

type Impl struct {
}

func (*Impl) Echo(ctx context.Context, in *EchoRequest) (*EchoResponse, error) {
	d := &EchoResponse{C: in.A, D: in.B}
	return d, nil
}

func (*Impl) SendCash() {
	fmt.Println("SendCash")
}

func (*Impl) GetOrder() {
	fmt.Println("GetOrder")
}

func main() {
	var c Test
	c = &Impl{}
	comm.ClientStub(c)
}
