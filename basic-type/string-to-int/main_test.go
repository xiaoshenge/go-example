package main

import (
	"testing"
)

func BenchmarkStringToIntWithAtoi(t *testing.B)  {
	for n := 0; n < t.N; n++ {
		StringToIntWithAtoi("219")
	}
}
func BenchmarkStringToIntWithSscanf(t *testing.B)  {
	for n := 0; n < t.N; n++ {
		StringToIntWithSscanf("219")
	}
}