package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func fetchContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func main()  {
	url := "https://example.com"
	content, err := fetchContent(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}