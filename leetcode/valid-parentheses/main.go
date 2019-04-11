package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Println(isValid("()[]{}"))
}
func isValid(s string) bool {
	var top int
	parentheses := make([]string,0)
	m := map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}
	for _, v := range s{
		v := string(v)
		var err error
		if v == "(" || v == "[" || v == "{" {
			parentheses = append(parentheses, v)
			top++
		} else if v == ")" || v == "]" || v == "}" {
			topStr := m[v]
			top, err = checkEnd(top, &parentheses, topStr)
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
	}
	if top == 0 {
		return true
	}
	return false
}
func checkEnd(top int, parentheses *[]string, topStr string) (int,  error) {
	if top < 1 {
		return top, errors.New("invalid1")
	}
	if ((*parentheses)[top-1] == topStr){
		*parentheses = (*parentheses)[0:top-1]
		top--
	} else {
		return top, errors.New("invalid2")
	}
	return top, nil
}