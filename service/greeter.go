package service

import (
	"context"
	greeter "plum/proto"
	"time"
)
type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	time.Sleep(2 * time.Second)
	return nil
}
