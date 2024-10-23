package katex

import (
	"io"
)

var inlineLeftB = []byte("\\(")
var inlineRightB = []byte("\\)")
var displayLeftB = []byte("\\[")
var displayRightB = []byte("\\]")

func Render(w io.Writer, src []byte, display bool) error {
	if display {
		w.Write(displayLeftB)
		w.Write(src)
		w.Write(displayRightB)
	} else {
		w.Write(inlineLeftB)
		w.Write(src)
		w.Write(inlineRightB)
	}
	return nil
}
