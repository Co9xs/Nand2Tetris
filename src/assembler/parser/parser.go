package parser

import (
	p "hack/common/parser"
	"strconv"
	"strings"
)

const (
	CmdTypeA = "a-instruction"
	CmdTypeC = "c-instruction"
	CmdTypeL = "label"
)

func New(sourceFile string) *p.Parser {
	return p.New(sourceFile)
}

func CommandType(c string) string {
	switch {
	case strings.HasPrefix(c, "(") && strings.HasSuffix(c, ")"):
		return CmdTypeL
	case strings.HasPrefix(c, "@"):
		return CmdTypeA
	default:
		return CmdTypeC
	}
}

func CommandArgs(s string) (a string, b string, c string) {
	a, b, c = "", "", ""
	switch CommandType(s) {
	case CmdTypeL:
		a = s[1 : len(s)-1]
	case CmdTypeA:
		a = s[1:]
	case CmdTypeC:
		compInd := strings.Index(s, "=")
		jumpInd := strings.Index(s, ";")
		if jumpInd != -1 {
			c = s[jumpInd+1:]
		} else {
			jumpInd = len(s)
		}
		if compInd == -1 {
			b = s[:jumpInd]
		} else {
			a = s[:compInd]
			b = s[compInd+1 : jumpInd]
		}
	}
	return
}

func IsVariable(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err != nil
}