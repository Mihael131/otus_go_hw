package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var letter rune
	var rB strings.Builder
	for _, r := range str {
		if unicode.IsDigit(r) {
			cnt, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			if letter == 0 || unicode.IsDigit(letter) {
				return "", ErrInvalidString
			}
			rB.WriteString(strings.Repeat(string(letter), cnt))
			letter = 0
		} else {
			if letter != 0 {
				rB.WriteRune(letter)
			}
			letter = r
		}
	}
	if letter != 0 {
		rB.WriteRune(letter)
	}
	return rB.String(), nil
}
