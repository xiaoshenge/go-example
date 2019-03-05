package main

import (
	"bytes"
	"os"
	"bufio"
	"fmt"
)

func main()  {
	fmt.Println("please enter some input:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var tmpbuf bytes.Buffer

	for scanner.Scan(){
		w := scanner.Text()
		tmpbuf.WriteString(w)
		if w[len(w)-1] == ';' && w[len(w)] == '\n'{
			fmt.Printf(tmpbuf.String())
			tmpbuf.Reset()
		} else {
			tmpbuf.WriteString(" ")
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}