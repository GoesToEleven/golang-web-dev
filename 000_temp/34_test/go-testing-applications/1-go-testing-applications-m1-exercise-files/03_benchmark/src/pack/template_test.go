package pack

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Hello, Go")

	b.ResetTimer()

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}
}
