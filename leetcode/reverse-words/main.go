package main

import (
	"fmt"
)

func reverseWords(s string) string {
	fmt.Println(s)
	str := ""
	tmp := ""
	for _,v := range s{
		fmt.Println(string(v))
		if string(v) != " "{
			tmp += string(v)
		} else {
			fmt.Println(string(tmp))
			str += string(reverString([]byte(tmp))) + " "
			tmp = ""
		}
	}
	if tmp!= ""{
		str += string(reverString([]byte(tmp)))
	}
	return str
}

func reverString(s []byte) []byte {
    len := len(s)
    if len <= 1 {
        return s
    }
    for i,j:=0,len/2;i<j;i++{
        k := len -i -1
        if i < k{
            s[i],s[k] = s[k],s[i]
        }
    }
    return s
}
func main()  {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}