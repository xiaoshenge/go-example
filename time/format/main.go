package main

import (
	"fmt"
	"time"
)

func main() {
	// // You can use Format argument of http://golang.org/pkg/time/#pkg-constants :
	// fmt.Println(time.Now().Format(time.Kitchen))

	// // also, you can input directly as string :
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println(time.Now().Format("15:04:05"))

	// // Timestamp
	// fmt.Println(time.Now().Unix()) // Ex: 1257894000

	// timeLocation, _ := time.LoadLocation("Local")
	// t := time.Date(2019, 2, 22, 12,0,0,0, timeLocation)
	// fmt.Println(t)

	// year,month,day := time.Now().Date()
	// fmt.Println(year)
	// fmt.Printf("%d\n",month)
	// fmt.Println(day)
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
	}
	time.Local = loc
	fmt.Println(time.Now().In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	test()

}
func test() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
