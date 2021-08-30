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

func PathInfo(path string) (name, dir string, isFile bool) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	splitted := strings.Split(path, "/")
	lastIndex := len(splitted) - 1
	nameExt := strings.Split(splitted[lastIndex], ".")
	name = nameExt[0]
	isDir := len(nameExt) == 1
	if isDir {
		dir = path + "/"
	} else {
		dir = strings.Join(splitted[:lastIndex], "/") + "/"
	}
	isFile = !isDir
	return
}

func ToXMLTag(tag, val string) string {
	return "<" + tag + ">" + normalizeXMLSymbol(val) + "</" + tag + ">"
}