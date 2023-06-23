package common

import (
	"errors"
	"strings"
	"unicode"
)

func CheckEmptyInput(input string) error {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 || isWhitespaceOnly(trimmed) {
		return errors.New(ErrorMissingInput)
	} else {
		return nil
	}
}

func isWhitespaceOnly(input string) bool {
	for _, r := range input {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
