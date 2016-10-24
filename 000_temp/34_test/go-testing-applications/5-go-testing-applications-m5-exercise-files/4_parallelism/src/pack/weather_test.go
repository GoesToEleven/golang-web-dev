package pack

import (
	"testing"
)

func BenchmarkPrintWeather(b *testing.B) {
	b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			PrintWeather(90001)
		}
	})
}
