package services

import "unicode"

func Capitalize(msg string) string {
	runes := []rune(msg)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
