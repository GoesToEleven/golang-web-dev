package pack

import (
	"testing"
)

func BenchmarkPrintWeather(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintWeather(90001)
	}
}
