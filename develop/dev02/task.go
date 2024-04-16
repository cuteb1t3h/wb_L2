package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	inputs := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "m2a3m6y8", "6e"}
	for _, input := range inputs {
		unpacked, err := stringUnpack(input)
		if err != nil {
			println("Error:", err.Error())
		} else {
			println("Unpacked:", unpacked)
		}
	}
}

func stringUnpack(s string) (string, error) {
	var result strings.Builder
	r := []rune(s)
	length := len(r)
	i := 0

	for i < length {
		cur := r[i]
		if cur == '\\' {
			if i+1 < length {
				next := r[i+1]
				if unicode.IsDigit(next) || next == '\\' {
					result.WriteRune(next)
					i += 2
					continue
				}
			}
			return "", errors.New("invalid string")
		}
		if unicode.IsDigit(cur) {
			return "", errors.New("invalid string")
		}
		if i+1 < length && unicode.IsDigit(r[i+1]) {
			count, err := strconv.Atoi(string(r[i+1]))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(string(cur), count))
			i += 2
		} else {
			result.WriteRune(cur)
			i++
		}
	}
	return result.String(), nil
}
