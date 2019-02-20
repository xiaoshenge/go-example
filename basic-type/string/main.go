package main

import (
	"fmt"
)

func StringJoinWithSprintf(s1,s2 string) string  {
	return fmt.Sprintf("%s %s", s1, s2)
}
func StringJoinWithAdd(s1,s2 string) string {
	return s1 + " " + s2
}