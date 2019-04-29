package main

import (
	"testing"
)

func BenchmarkSliceWithCopy(b *testing.B)  {
	for i:=0;i<b.N;i++{
		SliceWithCopy()
	}
}

func BenchmarkSliceWithSlice(b *testing.B)  {
	for i:=0;i<b.N;i++{
		SliceWithSlice()
	}
}