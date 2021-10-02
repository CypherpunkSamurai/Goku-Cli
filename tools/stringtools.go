package tools

import (
	"regexp"
)

func StripWhitespace(str string) string {
	/*
		Removes Extra Whitespaces
	*/
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, " ")
}
