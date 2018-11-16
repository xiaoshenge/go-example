package main

import (
    "bufio"
    "fmt"
    "net"
	"time"
	"example/tcp/protocol"
)

var quitSemaphore chan bool

func main()  {
	conn, err := net.Dial("tcp", ":8091")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	go func ()  {
		reader := bufio.NewReader(conn)
		for {
			// msg, err := reader.ReadString('<')
			// fmt.Println(msg)
			msg, _ := protocol.Decode(reader)
			fmt.Println(msg)
			if err != nil {
				quitSemaphore <- true
				break
			}
			time.Sleep(time.Second)
			// b := []byte("hi server, time:"+msg)
			// b := []byte("hi server>")
			b, _ := protocol.Encode("hi server")
			conn.Write(b)
		}
	}()
	// b := []byte("hi server>")
	b, _ := protocol.Encode("hi server")
    conn.Write(b)

    <-quitSemaphore
}