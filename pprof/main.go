package main

import (
	"log"
	"time"
	"runtime"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main()  {

	for index := 0; index < 100; index++ {
		go func ()  {
			time.Sleep(time.Second * 10000)
		}()
	}

	go func ()  {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	http.HandleFunc("/goroutinenum",func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, runtime.NumGoroutine())
	})
	err := http.ListenAndServe(":9091", nil)
	fmt.Println(err)
}