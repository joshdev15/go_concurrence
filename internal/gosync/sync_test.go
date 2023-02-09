package gosync

import "testing"

func Benchmark_secuential(b *testing.B) {
	for n := 0; n < b.N; n++ {
		secuential(urlList)
	}
}

func Benchmark_concurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concurrent(urlList)
	}
}
