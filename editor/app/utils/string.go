package utils

import (
	"fmt"
	"unicode/utf8"
)

func IntToRune(char int) rune {
	text := fmt.Sprint(char)
	r, _ := utf8.DecodeRuneInString(text)
	return r
}
