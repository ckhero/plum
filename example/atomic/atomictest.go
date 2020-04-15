package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Afuntion(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {


	// 输出结果：
	// finish
}

func add()  {
	var b int32 = 0
	var wg sync.WaitGroup
	var a int
	fmt.Println(wg)
	fmt.Println(a)
	//wg.Add(1000);
	ch := make(chan int, 1)
	<-ch
	for i := 0; i< 1000; i++ {
		go func() {
			atomic.AddInt32(&b, 1)
			//wg.Done()
		}()
	}
	ch<- 2
	//wg.Wait()
	fmt.Println(b)
}
