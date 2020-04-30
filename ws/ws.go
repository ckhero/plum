package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test1")

	//log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//http.ServeFile(w, r, "home.html")
}

func main() {
	//defer aaa()
	//int := 2
	//fmt.Println((int))
	//a := make(chan  bool)
	//b := make(chan  bool)
	//
	//	tick := time.NewTicker(5 * time.Second)
	//	go func(t *time.Ticker) {
	//		for {
	//			select {
	//			case <-t.C:
	//				fmt.Print("入队前的id：")
	//				return
	//			case <-a:
	//				fmt.Print("nextDone")
	//				return
	//			}
	//		}
	//	}(tick)
	//	time.Sleep(time.Second * 2)
	//	a <- true
	//	<-b

	a := make(chan bool)
	b := make(chan bool)
	 c := make(chan bool)
	go func() {
		for  {
			select {
			case <-a:
				fmt.Println("aaaaa")
			case <-b:
				fmt.Println("bbbbb")
			}
		}
	}()
	time.Sleep(time.Second * 1)
	go func() {
		a <- true
	}()
	go func() {
		b <- true
	}()

	time.Sleep(time.Second * 1)
	b <- true
	time.Sleep(time.Second * 1)
	a <- true
	time.Sleep(time.Second * 1)
	a <- true
	b <- true
	<- c
	//args := []int64{22222}
	//args = append(args, 333)
	//args = append(args, 333)
	//test22(args)
	//bindAddress := "localhost:2303"
	//r := gin.Default()
	//r.GET("/ping", test)
	//r.Run(bindAddress)
}
func aaa()  {
	fmt.Print(2222)

}
func test22(obj interface{})  {
	//var list []string

	//if reflect.TypeOf(obj).Kind() == reflect.Slice {
	//	s := reflect.ValueOf(obj)
	//	for i := 0; i < s.Len(); i++ {
	//		ele := s.Index(i)
	//		fmt.Print(ele.Int())
	//		fmt.Print(ele.Interface())
	//		//list = append(list, strconv.ParseInt())
	//	}
	//}
}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func test(c *gin.Context)  {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		fmt.Println(string(message))
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func test2(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("test2")
	fmt.Println(r.Body)
}