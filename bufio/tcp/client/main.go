package main

import (
    "fmt"
    "net"
)


func main()  {
	conn, err := net.Dial("tcp", ":8091")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

    conn.Write([]byte("set a 1\r#\nget a\r#\n"))

}