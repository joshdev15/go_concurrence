package channels

import "testing"

func Benchmark_signalExample(b *testing.B) {
	for n := 0; n < b.N; n++ {
		signalExample(urlList)
	}
}
