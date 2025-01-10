package benchmarking

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello-%d", i)
	}
}

// func BenchmarkBigLen(b *testing.B) {
// 	big := NewBig()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		big.Len()
// 	}
// }

// go test -cpu
func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}



// go test -benchmem -run=^$ -bench ^BenchmarkHello$