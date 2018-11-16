package main

import (
	// "time"
	"bufio"
	"fmt"
	"net"
	"example/tcp/protocol"
)

func main()  {
	l, err := net.Listen("tcp", ":8091")
	if err != nil{
		fmt.Println(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("connect from:", conn.RemoteAddr().String())
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn){

	defer func ()  {
		fmt.Println("disconnect by:", conn.RemoteAddr().String())
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

    for {
		// message, err := reader.ReadString('>')
		message, err := protocol.Decode(reader)
        if err != nil {
            return
        }

        fmt.Println(message)
		// msg := time.Now().String() + "<"
		// msg := "hi client<"
		// b := []byte(msg)
		b, _ := protocol.Encode("hi client")
        conn.Write(b)
    }
}