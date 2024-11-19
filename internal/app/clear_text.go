package app

import "regexp"

// clearText
func clearText(text string, reg *regexp.Regexp) string {
	return reg.ReplaceAllString(text, "")
}
