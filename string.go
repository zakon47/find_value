package find_value

import (
	"strings"
	"unicode/utf8"
)

//получить строковую подстроку слева => string,symbols
func String(str string) (string, string) {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return "", ""
	}

	var pos = 0
	for k := 0; k < len(str); k++ {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n <= 47 || n >= 58 {
			pos++
		} else {
			break
		}
	}
	var symbol, number string
	if pos == 0 {
		symbol = str[pos:]
		number = ""
	} else {
		symbol = str[pos:]
		number = str[:pos]
	}
	return number, symbol
}

//получить строковую подстроку справа => string,symbols
func StringReverse(str string) (string, string) {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return "", ""
	}

	var pos = len(str)
	for k := len(str) - 1; k >= 0; k-- {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n <= 47 || n >= 58 {
			pos--
		} else {
			break
		}
	}
	var symbol, number string
	if pos == len(str) {
		symbol = str[:pos]
		number = ""
	} else {
		symbol = str[:pos]
		number = str[pos:]
	}
	return number, symbol
}
