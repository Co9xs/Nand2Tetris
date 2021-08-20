package parser

import (
	"bufio"
	"hack/common/utils"
	"os"
	"strings"
)

const EOL = "\n"

type Parser struct {
	file *os.File
	scanner *bufio.Scanner
}

func New(path string) *Parser {
	f, err := os.Open(path)
	utils.HandleErr(err)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	return &Parser{f, scanner}
}

func (p *Parser) Parse() (string, bool) {
	s := ""
	ok := p.scanner.Scan()
	if !ok {
		return s, false
	}
	s = p.scanner.Text()
	if comment := strings.Index(s, "//"); comment > -1 {
		s = strings.TrimSpace(s[:comment])
	} else {
		s = strings.TrimSpace(s)
	}

	if len(s) == 0 {
		return p.Parse()
	}
	return s, true
}

func (p *Parser) Close() {
	p.file.Close()
}
