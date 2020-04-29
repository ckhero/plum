package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/micro/plugin"
	"io"
	"log"
	. "net/http"
	greeter "plum/proto"
	service2 "plum/service"
)

type Test struct {
	A string `json:"a"`
	BB Test2 `json:"bb"`
}
type Test2 struct {
	B string `json:"b"`
}
func main() {

	service := micro.NewService(
		micro.Name("greeter"),
	)
	plugin.Register(plugin.NewPlugin(
			plugin.WithName("breaker"),
			plugin.WithHandler(BreakerWrapper),
			plugin.WithInit(func(ctx *cli.Context) error {
			log.Println("Got value for example_flag", ctx.String("example_flag"))
			return nil
		}),
		))
	greeter.RegisterGreeterHandler(service.Server(), new(service2.Greeter))

	//srv.Handle("/", engine)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func BreakerWrapper(h Handler) Handler {
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		name := r.Method + "-" + r.RequestURI
		log.Println(name)
		err := hystrix.Do(name, func() error {
			h.ServeHTTP(w, r)

			sct := &StatusCodeTracker{ResponseWriter: w, status: StatusOK}
			h.ServeHTTP(sct.wrappedResponseWriter(), r)
			//
			if sct.status >= StatusBadRequest {
				str := fmt.Sprintf("status code %d", sct.status)
				log.Println(str)
				return errors.New(str)
			}

			if sct.status == 200 {
				str := fmt.Sprintf("status code %d", sct.status)
				log.Println(str)
				return errors.New(str)
			}
			return nil
		}, nil)
		if err != nil {
			log.Println("hystrix breaker err: ", err)
			return
		}
	})
}
type StatusCodeTracker struct {
	ResponseWriter
	status      int
	wroteheader bool
}

func (w *StatusCodeTracker) WriteHeader(status int) {
	w.status = status
	w.wroteheader = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *StatusCodeTracker) Write(b []byte) (int, error) {
	if !w.wroteheader {
		w.wroteheader = true
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}

// wrappedResponseWriter returns a wrapped version of the original
// ResponseWriter and only implements the same combination of additional
// interfaces as the original.  This implementation is based on
// https://github.com/felixge/httpsnoop.
func (w *StatusCodeTracker) wrappedResponseWriter()ResponseWriter {
	var (
		hj, i0 = w.ResponseWriter.(Hijacker)
		cn, i1 = w.ResponseWriter.(CloseNotifier)
		pu, i2 = w.ResponseWriter.(Pusher)
		fl, i3 = w.ResponseWriter.(Flusher)
		rf, i4 = w.ResponseWriter.(io.ReaderFrom)
	)

	switch {
	case !i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			ResponseWriter
		}{w}
	case !i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			ResponseWriter
			io.ReaderFrom
		}{w, rf}
	case !i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Flusher
		}{w, fl}
	case !i0 && !i1 && !i2 && i3 && i4:
		return struct {
			ResponseWriter
			Flusher
			io.ReaderFrom
		}{w, fl, rf}
	case !i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			Pusher
		}{w, pu}
	case !i0 && !i1 && i2 && !i3 && i4:
		return struct {
			ResponseWriter
			Pusher
			io.ReaderFrom
		}{w, pu, rf}
	case !i0 && !i1 && i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Pusher
			Flusher
		}{w, pu, fl}
	case !i0 && !i1 && i2 && i3 && i4:
		return struct {
			ResponseWriter
			Pusher
			Flusher
			io.ReaderFrom
		}{w, pu, fl, rf}
	case !i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			CloseNotifier
		}{w, cn}
	case !i0 && i1 && !i2 && !i3 && i4:
		return struct {
			ResponseWriter
			CloseNotifier
			io.ReaderFrom
		}{w, cn, rf}
	case !i0 && i1 && !i2 && i3 && !i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Flusher
		}{w, cn, fl}
	case !i0 && i1 && !i2 && i3 && i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Flusher
			io.ReaderFrom
		}{w, cn, fl, rf}
	case !i0 && i1 && i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Pusher
		}{w, cn, pu}
	case !i0 && i1 && i2 && !i3 && i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Pusher
			io.ReaderFrom
		}{w, cn, pu, rf}
	case !i0 && i1 && i2 && i3 && !i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Pusher
			Flusher
		}{w, cn, pu, fl}
	case !i0 && i1 && i2 && i3 && i4:
		return struct {
			ResponseWriter
			CloseNotifier
			Pusher
			Flusher
			io.ReaderFrom
		}{w, cn, pu, fl, rf}
	case i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
		}{w, hj}
	case i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			io.ReaderFrom
		}{w, hj, rf}
	case i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			Flusher
		}{w, hj, fl}
	case i0 && !i1 && !i2 && i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			Flusher
			io.ReaderFrom
		}{w, hj, fl, rf}
	case i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			Pusher
		}{w, hj, pu}
	case i0 && !i1 && i2 && !i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			Pusher
			io.ReaderFrom
		}{w, hj, pu, rf}
	case i0 && !i1 && i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			Pusher
			Flusher
		}{w, hj, pu, fl}
	case i0 && !i1 && i2 && i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			Pusher
			Flusher
			io.ReaderFrom
		}{w, hj, pu, fl, rf}
	case i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
		}{w, hj, cn}
	case i0 && i1 && !i2 && !i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			io.ReaderFrom
		}{w, hj, cn, rf}
	case i0 && i1 && !i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Flusher
		}{w, hj, cn, fl}
	case i0 && i1 && !i2 && i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Flusher
			io.ReaderFrom
		}{w, hj, cn, fl, rf}
	case i0 && i1 && i2 && !i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Pusher
		}{w, hj, cn, pu}
	case i0 && i1 && i2 && !i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Pusher
			io.ReaderFrom
		}{w, hj, cn, pu, rf}
	case i0 && i1 && i2 && i3 && !i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Pusher
			Flusher
		}{w, hj, cn, pu, fl}
	case i0 && i1 && i2 && i3 && i4:
		return struct {
			ResponseWriter
			Hijacker
			CloseNotifier
			Pusher
			Flusher
			io.ReaderFrom
		}{w, hj, cn, pu, fl, rf}
	default:
		return struct {
			ResponseWriter
		}{w}
	}
}

