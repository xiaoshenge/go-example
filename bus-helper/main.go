package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", getBus)
	http.ListenAndServe(":9091", nil)
}

func getBus(w http.ResponseWriter, r *http.Request) {
	busCh := make(chan string, 4)
	busList := []Bus{
		Bus{Title: "松江66路", Data: url.Values{"stoptype": {"1"}, "stopid": {"24."}, "sid": {"e78627ed5d1a8b81728da2bf62522857"}}},
		Bus{Title: "松江21路", Data: url.Values{"stoptype": {"1"}, "stopid": {"17."}, "sid": {"ba7da5433ee494fa042143e4658093ae"}}},
		Bus{Title: "松江7路", Data: url.Values{"stoptype": {"0"}, "stopid": {"24."}, "sid": {"4558c1636efdee980c10f788ddf2b37a"}}},
		Bus{Title: "松江64路", Data: url.Values{"stoptype": {"1"}, "stopid": {"6."}, "sid": {"43b249e9065360e2f3dbddd33836eb8c"}}},
	}

	for _, bus := range busList {
		go func(b Bus) {
			busCh <- getBusInfoWithTimeout(b.Title, b.Data)
		}(bus)
	}

	var res string
	for index := 0; index < 4; index++ {
		res += <-busCh + "\n"
	}
	fmt.Fprint(w, res)
}

type Bus struct {
	Title string
	Data  url.Values
}

func getBusInfoWithTimeout(busNo string, data url.Values) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	ch := make(chan string)

	go func() {
		getBusInfo(busNo, data, ch)
	}()

	select {
	case <-ctx.Done():
		return "timeout"
	case str := <-ch:
		return str
	}
}

func getBusInfo(busNo string, data url.Values, ch chan string) {
	resp, err := http.PostForm("https://shanghaicity.openservice.kankanews.com/public/bus/Getstop", data)
	if err != nil {
		ch <- busNo + " 请求异常"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- busNo + " 请求异常"
	}

	res := BusPos{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		ch <- busNo + " 待发车"
	}

	if len(res) != 1 {
		ch <- busNo + " 待发车"
	}

	fmt.Println(res)

	sec, err := strconv.ParseFloat(res[0].Time, 32)
	if err != nil {
		ch <- busNo + " 待发车"
	}
	min := math.Ceil(sec / 60.0)

	ch <- busNo + " 还有" + res[0].Stopdis + "站" + " 约" + fmt.Sprintf("%.0f", min) + "分钟"
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
