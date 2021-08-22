package codewriter

import (
	"fmt"
	"hack/common/codewriter"
	"hack/common/writer"
	"strconv"
	"strings"
)

type CodeWriter struct {
	*writer.Writer
	count int
	fileName
}