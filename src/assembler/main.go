package assembler

import (
	"hack/assembler/codewriter"
	"hack/assembler/parser"
	"hack/assembler/symboltable"
	"hack/common/utils"
	"os"
	"strconv"
)

func main() {
	Compile(os.Args[1], os.Args[2])
}

func Compile(src string, out string) {
	var lines []string
	st := symboltable.New()
	p := parser.New(src)
	cw := codewriter.New(out)
	defer p.Close()
	defer cw.Close()

	hasMore := true
	for hasMore {
		c, ok := p.Parse()
		hasMore = ok
		if ok {
			if parser.CommandType(c) != parser.CmdTypeL {
				lines = append(lines, c)
			} else {
				label, _, _ := parser.CommandArgs(c)
				st[label] = len(lines)
			}
		}
	}

	n := 16
	for _, l := range lines {
		cmdType := parser.CommandType(l)
		arg1, arg2, arg3 := parser.CommandArgs(l)

		switch cmdType {
		case parser.CmdTypeA:
			if parser.IsVariable(arg1) {
				_, found := st[arg1]
				if !found {
					st[arg1] = n
					n++
				}
				cw.WriteA(st[arg1])
			} else {
				aInt, err := strconv.Atoi(arg1)
				utils.HandleErr(err)
				cw.WriteA(aInt)
			}
		case parser.CmdTypeC:
			cw.WriteC(arg1, arg2, arg3)
		}
	}

	utils.LogDone(src, out)
}