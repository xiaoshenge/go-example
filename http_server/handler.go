package main

import "net/http"

func main() {
	mux := &cusHandler{}
	http.ListenAndServe(":8080", mux)
}

type cusHandler struct {

}

func (h *cusHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path == "/test"{
		w.Write([]byte("hello test"))
		return
	}

	http.NotFound(w, r)
}
