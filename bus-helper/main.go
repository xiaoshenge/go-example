package main

import (
	"math"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", getBus)
	http.ListenAndServe(":9091", nil)
}

func getBus(w http.ResponseWriter, r *http.Request)  {
	res := getBusInfo("松江66路", url.Values{"stoptype":{"1"}, "stopid":{"24."},"sid":{"e78627ed5d1a8b81728da2bf62522857"}}) + "\n" + 
	getBusInfo("松江21路", url.Values{"stoptype":{"1"}, "stopid":{"17."},"sid":{"ba7da5433ee494fa042143e4658093ae"}}) + "\n" + 
	getBusInfo("松江7路", url.Values{"stoptype":{"0"}, "stopid":{"24."},"sid":{"4558c1636efdee980c10f788ddf2b37a"}}) + "\n" + 
	getBusInfo("松江64路", url.Values{"stoptype":{"1"}, "stopid":{"6."},"sid":{"43b249e9065360e2f3dbddd33836eb8c"}}) + "\n" 
	
	fmt.Fprint(w, res)
}


func getBusInfo(busNo string, data url.Values) string {
	resp, err := http.PostForm("https://shanghaicity.openservice.kankanews.com/public/bus/Getstop", data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return busNo + " 请求异常"
	}


	res := BusPos{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return busNo + " 待发车"
	}

	sec,err  := strconv.ParseFloat(res[0].Time, 32)
	if err != nil {
		return busNo + " 待发车"
	}
	min := math.Ceil(sec/60.0)
	
	return busNo + " 还有" + res[0].Stopdis + "站" + " 约" + fmt.Sprintf("%.0f", min) + "分钟"
}

type BusPos []struct {
	Attributes struct {
		Cod string `json:"cod"`
	} `json:"@attributes"`
	Terminal string `json:"terminal"`
	Stopdis  string `json:"stopdis"`
	Distance string `json:"distance"`
	Time     string `json:"time"`
}