package service

import (
	"context"
	test "plum/proto/test"
	"time"
)

type Test struct {}

func (g *Test) Hello(ctx context.Context, req *test.HelloRequest, rsp *test.HelloResponse) error {
	time.Sleep(2 * time.Second)
	return nil
}
