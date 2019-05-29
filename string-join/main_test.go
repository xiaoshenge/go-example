package join

import "testing"

const BLOG = "http://www.golang.org/"

func initStrings(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func initStringi(N int) []interface{} {
	s := make([]interface{}, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func BenchmarkStringPlus1000(b *testing.B) {
	p := initStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

// func BenchmarkStringFmt1000(b *testing.B) {
// 	p := initStringi(1000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		StringFmt(p)
// 	}
// }

// func BenchmarkStringJoin1000(b *testing.B) {
// 	p := initStrings(1000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		StringJoin(p)
// 	}
// }

// func BenchmarkStringBuffer1000(b *testing.B) {
// 	p := initStrings(1000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		StringBuffer(p)
// 	}
// }

// func BenchmarkStringBuilder1000(b *testing.B) {
// 	p := initStrings(1000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		StringBuilder(p)
// 	}
// }
