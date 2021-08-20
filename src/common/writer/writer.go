package writer

import (
	"hack/common/utils"
	"os"
)

type Writer struct {
	file *os.File
	WriterInterface
}

type WriterInterface interface {
	Close()
	WriteLine(s string)
}

func New(outFile string) *Writer {
	f, err := os.Create(outFile)
	utils.HandleErr(err)
	return &Writer{
		file: f,
	}
}

func (w *Writer) Close() {
	w.file.Close()
}

func (w *Writer) WriteLine(s string) {
	fi, err := w.file.Stat()
	utils.HandleErr(err)
	eol := "\n"
	if fi.Size() == 0 {
		eol = ""
	}
	w.file.WriteString(eol + s)
}