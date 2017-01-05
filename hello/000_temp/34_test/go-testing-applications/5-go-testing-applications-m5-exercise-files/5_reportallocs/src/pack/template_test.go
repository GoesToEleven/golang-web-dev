package pack

import (
	"testing"
)

func BenchmarkTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := NewUser("Arthur")
		SayHello(user)
	}
}

func BenchmarkTemplateCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := NewUser("Arthur")
		SayHello(user)
	}
}
