package main

import (
	"log"
	"time"
	"context"
	"fmt"
	"net/http"
)

func main()  {
	uri := "https://httpbin.org/delay/3"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond * 100)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()
}