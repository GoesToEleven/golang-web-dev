package pack

import (
	"testing"
)

func BenchmarkTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		user := NewUser("Arthur")
		b.StartTimer()
		SayHello(user)
	}
}
