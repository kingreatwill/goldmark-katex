# goldmark-katex


Copy for [goldmark-mathjax](https://github.com/litao91/goldmark-mathjax).
Used by [mdserver](https://github.com/kingreatwill/mdserver).

## Installation


```
go get github.com/kingreatwill/goldmark-katex
```

## Usage

```go
package main

import (
	"bytes"
	"fmt"

	katex "github.com/kingreatwill/goldmark-katex"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(katex.KaTeX),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	// todo more control on the parsing process
	var html bytes.Buffer
	mdContent := []byte(`
$$
\mathbb{E}(X) = \int x d F(x) = \left\{ \begin{aligned} \sum_x x f(x) \; & \text{ if } X \text{ is discrete} 
\\ \int x f(x) dx \; & \text{ if } X \text{ is continuous }
\end{aligned} \right.
$$


Inline math $\frac{1}{2}$
`)
	if err := md.Convert(mdContent, &html); err != nil {
		fmt.Println(err)
	}
	fmt.Println(html.String())
}
```

License
--------------------
MIT

