package main

import "fmt"

func main() {
	phone := "13564663499"
	if len(phone) != 11 {
		fmt.Println("invalid")
	}

	str := fmt.Sprintf("%s****%s", phone[0:3], phone[7:11])

	fmt.Println(str)
}
