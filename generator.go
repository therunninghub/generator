package generator

import (
	"errors"
	"math/rand"
	"strings"
)

// Charsets
const (
	Uppercase         string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase         string = "abcdefghijklmnopqrstuvwxyz"
	Numeric           string = "0123456789"
	Alphabetic        string = Uppercase + Lowercase
	AlphaNumeric      string = Alphabetic + Numeric
	UpperAlphaNumeric string = Uppercase + Numeric
	LowerAlphaNumeric string = Lowercase + Numeric
)

// String Return a random string with input length
func String(length int) (string, error) {
	return StringWithCharset(length, AlphaNumeric)
}

// StringWithCharset Return a random string with input length and specific charset
func StringWithCharset(length int, charset string) (string, error) {
	if length < 1 {
		return "", errors.New("Invalid input length")
	}

	_charset := charset
	if charset == "" {
		_charset = AlphaNumeric
	}

	charsetLength := len(_charset)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(charsetLength)])
	}

	return sb.String(), nil
}

// StringListWithCharset Return a list of random string with input length and specific charset and input number of strings to be generated
func StringListWithCharset(length int, charset string, number int) ([]string, error) {
	if number < 1 {
		return nil, errors.New("Invalid input number")
	}

	exist := make(map[string]bool)
	list := make([]string, number)
	i := 0
	for i < number {
		s, err := StringWithCharset(length, charset)
		if err != nil {
			return nil, err
		}

		if len(s) == 0 {
			continue
		}

		exists := exist[s]
		if !exists {
			exist[s] = true
			list[i] = s
			i++
		}

	}

	return list, nil
}
