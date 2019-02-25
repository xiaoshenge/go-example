package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
)


func Logger(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(w, r)
        endTime := time.Since(startTime)
        log.Printf("%s %d %v", r.URL, r.Method, endTime)
    })
}

func loginHanlder(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("in hanlder func")
}

func main() {
	// http.handler("/login", Logger(loginHanlder))
	http.Handle("/login", Logger(http.HandlerFunc(loginHanlder)))
    http.ListenAndServe(":9091", nil)
}