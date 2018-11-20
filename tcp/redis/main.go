package main

import (
	"bytes"
	"bufio"
	"fmt"
	"net"
)

func main()  {
	l, err := net.Listen("tcp", ":6380")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil{
			fmt.Println("err")
		}
		go handlerReq(conn)
	}
}

func handlerReq(conn net.Conn)  {
	defer func ()  {
		conn.Close()
	}()


	// scanner := bufio.NewScanner(conn)
	// scanner.Split(ScanCRLF)

	// for scanner.Scan() {
	// 	fmt.Println("input: ", scanner.Text())
	// 	conn.Write([]byte("+OK\r\n"))
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Printf("Invalid input: %s", err)
	// }

	reader := bufio.NewReader(conn)
	for {
		// msg, prefix, err := reader.ReadLine()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println("prefix:", prefix)
		// fmt.Println("msg: ",string(msg))
		// conn.Write([]byte("+OK\r\n"))

		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(msg)
		conn.Write([]byte("+OK\r\n"))
	}

}

func dropCR(data []byte) []byte{
	if len(data) > 0 && data[len(data) -1] == '\r'{
		return data[0:len(data)-1]
	}
	return data
}
func ScanCRLF(data []byte, atEOF bool)(advance int ,token []byte, err error)  {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0{
		return i + 2, dropCR(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}