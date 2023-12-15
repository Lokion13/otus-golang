package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) {
	if len(inputStr) == 0 {
		return "", nil
	}
	_, err := strconv.Atoi(string([]rune(inputStr)[0]))
	if err == nil {
		return "", ErrInvalidString
	}

	resStr := strings.Builder{}

	for i, v := range inputStr {
		rInt, ErrIsDigit := strconv.Atoi(string(inputStr[i]))

		switch {
		case ErrIsDigit == nil && rInt >= 0 && rInt < 10:
			if rInt == 0 {
				str := resStr.String()
				str = str[:len(str)-1]
				resStr.Reset()
				resStr.WriteString(str)
			} else {
				resStr.WriteString(strings.Repeat(string(inputStr[i-1]), rInt-1))
			}
		default:
			resStr.WriteString(string(v))
		}
	}
	return resStr.String(), nil
}
