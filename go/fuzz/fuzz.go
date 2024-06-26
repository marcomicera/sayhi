package fuzz

import (
	"fmt"
	su "github.com/marcomicera/sayhi/go/stringutils"
	"strings"
	"unicode"
)

func RemoveWhitespaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func Fuzz(data []byte) int {

	name := string(data)
	splitCamelCase, err := su.SplitCamelCasePersonName(name)

	// Ignoring names with invalid characters
	if err != nil {
		return -1
	}

	revertedNameWithoutSpaces := RemoveWhitespaces(splitCamelCase)
	if name != revertedNameWithoutSpaces {
		panic(fmt.Sprintf("Expected %q, got %q", name, revertedNameWithoutSpaces))
	}
	return 1
}
