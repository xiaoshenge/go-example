package main

import "net/http"

func main() {
	http.HandleFunc("/test", testHandler)

	http.ListenAndServe(":8080", nil)
}

func testHandler(writer http.ResponseWriter, req *http.Request)  {
	writer.Write([]byte("hello test"))
}
