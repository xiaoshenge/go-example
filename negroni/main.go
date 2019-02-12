package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/urfave/negroni"
)

func main()  {
	n := negroni.Classic()
	n.UseFunc(testMiddle1)
	n.UseFunc(testMiddle2)
	n.UseHandler(handler())
	n.Run(":1234")
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler)
}

func myHandler(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "hello world")
}

func testMiddle1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("middle1 start")
	next(rw, r)
	fmt.Println("middle1 end")
}
func testMiddle2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("middle2 start")
	next(rw, r)
	fmt.Println("middle2 end")
}