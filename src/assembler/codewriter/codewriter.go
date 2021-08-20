package codewriter

import (
	"fmt"
	"hack/common/writer"
)

type CodeWriter struct {
	*writer.Writer
}

func New(outFile string) *CodeWriter {
	return &CodeWriter{writer.New(outFile)}
}

func (cw *CodeWriter) WriteA(v int) {
	s := fmt.Sprintf("%015b", v)
	hack := "0" + s[len(s)-15:]
	cw.WriteLine(hack)
}

func (cw *CodeWriter) WriteC(d, c, j string) {
	if len(d) == 0 {
		d = "null"
	}
	if len(j) == 0 {
		j = "null"
	}
	hack := fmt.Sprintf("111%s%s%s", computations[c], destinations[d], jumps[j])
	cw.WriteLine(hack)
}