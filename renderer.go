package katex

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type HTMLRenderer struct {
	html.Config
}

func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindInline, r.renderInline)
	reg.Register(KindBlock, r.renderBlock)
}

func (r *HTMLRenderer) renderInline(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		node := n.(*Inline)

		b := bytes.Buffer{}
		err := Render(&b, node.Equation, false)
		if err != nil {
			return ast.WalkStop, err
		}
		html := b.Bytes()
		w.Write(html)
		return ast.WalkContinue, nil
	}

	return ast.WalkContinue, nil
}

func (r *HTMLRenderer) renderBlock(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		node := n.(*Block)

		b := bytes.Buffer{}
		err := Render(&b, node.Equation, true)
		if err != nil {
			return ast.WalkStop, err
		}
		html := b.Bytes()
		w.WriteString("<div>")
		w.Write(html)
		w.WriteString("</div>")
		return ast.WalkContinue, nil
	}

	return ast.WalkContinue, nil
}
