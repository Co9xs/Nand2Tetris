package utils

import (
	"fmt"
	"strings"
)

func HandleErr(e error) {
	if e != nil {
		panic(e)
	}
}