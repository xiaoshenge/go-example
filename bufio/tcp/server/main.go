package main

import (
	"strings"
	// "time"
	"bufio"
	"fmt"
	"net"
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
	
	scanner := bufio.NewScanner(conn)
	scanner.Split(splitFunc)

	// var tmpbuf bytes.Buffer

	for scanner.Scan(){
		w := scanner.Text()
		fmt.Printf("%#v\n", w)
	}
}
func splitFunc(data []byte, atEOF bool)(advance int, token []byte, err error)  {
	if atEOF && len(data) == 0{
		return 0, nil, nil
	}

	if i:=strings.Index(string(data), "\r#\n"); i >= 0 {
		return i+3, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}