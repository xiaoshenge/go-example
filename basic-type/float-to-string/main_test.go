package main

import (
	"testing"
)

func BenchmarkFloatToStringWithFromatFloat(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		FloatToStringWithFromatFloat(3.1415926)
	}
}
func BenchmarkFloatToStringWithSprintf(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		FloatToStringWithSprintf(3.1415926)
	}
}