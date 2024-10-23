package katex

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleRender() {
	b := bytes.Buffer{}
	Render(&b, []byte(`Y = A \dot X^2 + B \dot X + C`), false)
	fmt.Println(b.String())
}

func BenchmarkRender(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := bytes.Buffer{}
		Render(&b, []byte(`Y = A \dot X^2 + B \dot X + C`), false)
	}
}
