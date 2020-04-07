package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"net"
	. "net/http"
	greeter "plum/proto"
	"time"
)

func main() {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)
	service := micro.NewService(
		micro.Name("greete22r"),
		micro.WrapClient(NewProdsWrapper),
	)
	//plugin.Register(plugin.NewPlugin(
	//		plugin.WithName("breaker"),
	//		plugin.WithHandler(BreakerWrapper),
	//		plugin.WithInit(func(ctx *cli.Context) error {
	//		log.Println("Got value for example_flag", ctx.String("example_flag"))
	//		return nil
	//	}),
	//	))
	service.Init()
	greeter22 := greeter.NewGreeterService("greeter", service.Client())

	// request the Hello method on the Greeter handler
	for i := 0; i< 10; i++ {
		go greeter22.Hello(context.TODO(), &greeter.HelloRequest{
			Name:                 "ck",
		})
		go greeter22.Hello(context.TODO(), &greeter.HelloRequest{
			Name:                 "ck",
		})
			rsp, err := greeter22.Hello(context.TODO(), &greeter.HelloRequest{
				Name:                 "ck",
			})

			if i == 5 {
				time.Sleep(2 * time.Second)
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rsp.Greeting)
	}
}

//func BreakerWrapper(h Handler) Handler {
//	return HandlerFunc(func(w ResponseWriter, r *Request) {
//		name := r.Method + "-" + r.RequestURI
//		log.Println(name)
//		err := hystrix.Do(name, func() error {
//			h.ServeHTTP(w, r)
//
//			sct := &statusCodeTracker{ResponseWriter: w, status: StatusOK}
//			h.ServeHTTP(sct.wrappedResponseWriter(), r)
//			//
//			if sct.status >= StatusBadRequest {
//				str := fmt.Sprintf("status code %d", sct.status)
//				log.Println(str)
//				return errors.New(str)
//			}
//
//			if sct.status == 200 {
//				str := fmt.Sprintf("status code %d", sct.status)
//				log.Println(str)
//				return errors.New(str)
//			}
//			return nil
//		}, nil)
//		if err != nil {
//			log.Println("hystrix breaker err: ", err)
//			return
//		}
//	})
//}
