package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"github.com/urfave/negroni"
	"fmt"
)

func main() {
	router := httprouter.New()

	router.POST("/test", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("test"))
	})

	router.POST("/test1", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("test1"))
	})

	n := negroni.New()
	n.Use(negroni.HandlerFunc(MyMiddleware3))
	n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("a")
		rw.Write([]byte("a"))
		next(rw,r)
	})
	n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Println("aaa")
		rw.Write([]byte("aa"))
		next(rw,r)
	})
	router.Handler(http.MethodPost, "/a", n)


	h := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 2 << 15,
	}

	h.ListenAndServe()

}

func MyMiddleware3(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("myMiddleware3")
	next(rw, r)
	fmt.Println("myMiddleware3 end")
}