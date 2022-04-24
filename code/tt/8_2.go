package main

//package tt
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//	_ "net/http/pprof"
//)
//
//func ll() {
//	fmt.Println("123")
//
//}
//
//func main() {
//	go func() {
//		log.Println(http.ListenAndServe("localhost:9981", nil))
//	}()
//	for i := 0; i < 10000000000; i++ {
//		ll()
//	}
//	//time.Sleep(3 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//	//time.Sleep(1 * time.Second)
//	//ll()
//}

import (
	"sync"
	"time"
)

var m = make(map[int]int)

func main() {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		for {
			w := make(map[int]int)
			m = w
		}
	}()

	go func() {
		wg.Add(1)
		for {
			for _, v := range m {
				_ = v
			}
		}
	}()
	wg.Wait()
	time.Sleep(10 * time.Second)
}
