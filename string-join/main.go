package join

import (
	"bytes"
	"fmt"
	"strings"
)

func StringPlus(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

func StringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}

func StringJoin(p []string) string {
	return strings.Join(p, "")
}

func StringBuffer(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
