package stringutils

import (
	"errors"
	"regexp"
	"strings"
)

// Given a camel case string, it splits it using whitespaces
func SplitCamelCasePersonName(personName string) (string, error) {

	// Checking input
	personNameRegex, _ := regexp.Compile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
	if !personNameRegex.MatchString(personName) {
		return "", errors.New("invalid person name")
	}

	var a []string
	var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)") // camel regex
	for _, sub := range camel.FindAllStringSubmatch(personName, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.Join(a, " "), nil
}
