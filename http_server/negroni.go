package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
	)

func main() {


	n := negroni.New()

	n.Use(negroni.HandlerFunc(MyMiddleware))
	n.UseFunc(negroni.HandlerFunc(MyMiddleware2))

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test"))
		fmt.Println("in test path")
	})
	mux.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test1"))
		fmt.Println("in test1 path")
	})


	n.UseHandler(mux)


	//n.Run(":8081")

	http.ListenAndServe(":8081",n)
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("myMiddleware1")
	next(rw, r)
}

func MyMiddleware2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("myMiddleware2")
	next(rw, r)
}