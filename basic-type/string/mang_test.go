package main

import (
	"testing"
)

func BenchmarkStringJoinWithSprintf(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		StringJoinWithSprintf("hello", "world")
	}
}
func BenchmarkStringJoinWithAdd(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		StringJoinWithAdd("hello", "world")
	}
}