package katex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type mathJaxBlockParser struct {
}

var defaultMathJaxBlockParser = &mathJaxBlockParser{}

type mathBlockData struct {
	indent int
}

var mathBlockInfoKey = parser.NewContextKey()

func NewMathJaxBlockParser() parser.BlockParser {
	return defaultMathJaxBlockParser
}

func (b *mathJaxBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos == -1 {
		return nil, parser.NoChildren
	}
	if line[pos] != '$' {
		return nil, parser.NoChildren
	}
	i := pos
	for ; i < len(line) && line[i] == '$'; i++ {
	}
	if i-pos < 2 {
		return nil, parser.NoChildren
	}
	pc.Set(mathBlockInfoKey, &mathBlockData{indent: pos})
	node := NewMathBlock()
	return node, parser.NoChildren
}

func (b *mathJaxBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()
	data := pc.Get(mathBlockInfoKey).(*mathBlockData)
	w, pos := util.IndentWidth(line, 0)
	if w < 4 {
		i := pos
		for ; i < len(line) && line[i] == '$'; i++ {
		}
		length := i - pos
		if length >= 2 && util.IsBlank(line[i:]) {
			reader.Advance(segment.Stop - segment.Start - segment.Padding)
			return parser.Close
		}
	}

	pos, padding := DedentPosition(line, 0, data.indent)
	seg := text.NewSegmentPadding(segment.Start+pos, segment.Stop, padding)
	node.Lines().Append(seg)
	reader.AdvanceAndSetPadding(segment.Stop-segment.Start-pos-1, padding)
	return parser.Continue | parser.NoChildren
}

// DedentPosition dedents lines by the given width.
func DedentPosition(bs []byte, currentPos, width int) (pos, padding int) {
	if width == 0 {
		return 0, 0
	}
	w := 0
	l := len(bs)
	i := 0
	for ; i < l; i++ {
		if bs[i] == '\t' {
			w += util.TabWidth(currentPos + w)
		} else if bs[i] == ' ' {
			w++
		} else {
			break
		}
	}
	if w >= width {
		return i, w - width
	}
	return i, 0
}

func (b *mathJaxBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	pc.Set(mathBlockInfoKey, nil)
}

func (b *mathJaxBlockParser) CanInterruptParagraph() bool {
	return true
}

func (b *mathJaxBlockParser) CanAcceptIndentedLine() bool {
	return false
}

func (b *mathJaxBlockParser) Trigger() []byte {
	return nil
}
