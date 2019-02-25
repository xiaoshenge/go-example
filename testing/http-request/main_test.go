package main

import (
	"net/http"
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_fetchContent(t *testing.T)  {
	ts := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "hello world")
	}))	
	defer ts.Close()

	content, err := fetchContent(ts.URL)
	if err != nil {
		t.Error(err)
	}
	
	want := "hello world"
	if content != want {
		t.Errorf("Got %q, want %q", content, want)
	}
}