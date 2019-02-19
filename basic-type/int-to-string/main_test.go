package main

import (
	"testing"
)

func BenchmarkIntToStingWithItoa(t *testing.B)  {
	for n := 0; n < t.N; n++ {
		IntToStingWithItoa(219)
	}
}

func BenchmarkIntToStringWithSprinf(t *testing.B)  {
	for n := 0; n < t.N; n++ {
		IntToStringWithSprinf(129)
	}
}