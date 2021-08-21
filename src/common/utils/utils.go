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

func LogDone(a string, b string) {
	fmt.Printf("- %-30s -> %-30s \xE2\x9C\x94 done\n", a, b)
}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}

func CompareStrings(a, b string) bool {
	return filterNewLines(a) == filterNewLines(b)
}