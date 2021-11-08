package katex

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/yuin/goldmark"

	"github.com/stretchr/testify/assert"
)

type katexTestCase struct {
	d   string // test description
	in  string // input markdown source
	out string // expected output html
}

func TestKaTeX(t *testing.T) {

	tests := []katexTestCase{
		{
			d:   "plain text",
			in:  "foo",
			out: `<p>foo</p>`,
		},
		{
			d:   "bold",
			in:  "**foo**",
			out: `<p><strong>foo</strong></p>`,
		},
		{
			d:   "math inline",
			in:  "$1+2$",
			out: `<p><span class="math inline">\(1+2\)</span></p>`,
		},
		{
			d:  "math display",
			in: "$$\n1+2\n$$",
			out: `<p><span class="math display">\[1+2
\]</span></p>`,
		},
		{
			// this input previously triggered a panic in block.go
			d:   "list-begin",
			in:  "*foo\n  ",
			out: "<p>*foo</p>",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tc.d), func(t *testing.T) {
			out, err := renderMarkdown([]byte(tc.in))
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.out, strings.TrimSpace(string(out)))
		})
	}

}

func renderMarkdown(src []byte) ([]byte, error) {
	md := goldmark.New(
		goldmark.WithExtensions(KaTeX),
	)

	var buf bytes.Buffer
	if err := md.Convert(src, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
