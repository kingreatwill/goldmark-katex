package katex

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type katex struct {
	inlineStartDelim string
	inlineEndDelim   string
	blockStartDelim  string
	blockEndDelim    string
}

type Option interface {
	SetOption(e *katex)
}

type withInlineDelim struct {
	start string
	end   string
}

type withBlockDelim struct {
	start string
	end   string
}

func WithInlineDelim(start string, end string) Option {
	return &withInlineDelim{start, end}
}

func (o *withInlineDelim) SetOption(e *katex) {
	e.inlineStartDelim = o.start
	e.inlineEndDelim = o.end
}

func WithBlockDelim(start string, end string) Option {
	return &withBlockDelim{start, end}
}

func (o *withBlockDelim) SetOption(e *katex) {
	e.blockStartDelim = o.start
	e.blockEndDelim = o.end
}

var KaTeX = &katex{
	inlineStartDelim: `\(`,
	inlineEndDelim:   `\)`,
	blockStartDelim:  `\[`,
	blockEndDelim:    `\]`,
}

func NewKaTeX(opts ...Option) *katex {
	r := &katex{
		inlineStartDelim: `\(`,
		inlineEndDelim:   `\)`,
		blockStartDelim:  `\[`,
		blockEndDelim:    `\]`,
	}

	for _, o := range opts {
		o.SetOption(r)
	}
	return r
}

func (e *katex) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewMathJaxBlockParser(), 701),
	))
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewInlineMathParser(), 501),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewMathBlockRenderer(e.blockStartDelim, e.blockEndDelim), 501),
		util.Prioritized(NewInlineMathRenderer(e.inlineStartDelim, e.inlineEndDelim), 502),
	))
}
