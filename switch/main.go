package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("i", 0, "int")
	flag.Parse()
loop:
	for j := 0; j < 3; j++ {
		switch *i {
		case 0, 1, 2, 3:
			fmt.Println("less then 4")
			break loop
		case 4, 5:
			fmt.Println(">4 , 5")
		}
	}

}
