package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("Введите строку:")
	var (
		str string
	)
	fmt.Scan(&str)

	slise := []rune(str)

	if strings.Contains(str, "\\") {
		fmt.Println(stringBuilderWithEscape(slise))
	} else {
		fmt.Println(stringBuilder(slise))
	}
}

func stringBuilder(slise []rune) (string, error) {
	res := ""
	for i := 0; i < len(slise); i++ {
		compare := slise[i]
		if unicode.IsDigit(compare) {
			if i == 0 || unicode.IsDigit(slise[i+1]) {
				err := fmt.Errorf("(некорректная строка)")
				return "", err
			}
			count, _ := strconv.Atoi(string(compare))
			addstr := ""
			for j := 0; j < count-1; j++ {
				addstr += string(slise[i-1])
			}
			if addstr != "" {
				res += addstr
			}
			continue
		}
		res += string(slise[i])
	}
	return res, nil
}

func stringBuilderWithEscape(slice []rune) (string, error) {
	res := ""
	escaping := false
	escapeValue := ""
	for i := 0; i < len(slice); i++ {
		compare := slice[i]
		if escaping {
			escapeValue += string(compare)
			escaping = false
		} else if compare == '\\' {
			escaping = true
		} else if unicode.IsDigit(compare) {
			if escapeValue != "" {
				count, _ := strconv.Atoi(string(compare))
				addStr := ""
				for j := 0; j < count; j++ {
					addStr += escapeValue
				}
				if addStr != "" {
					res += addStr
				}
				escapeValue = ""
			} else {
				if i == 0 || unicode.IsDigit(slice[i-1]) {
					err := fmt.Errorf("(некорректная строка)")
					return "", err
				}
				count, _ := strconv.Atoi(string(compare))
				addStr := ""
				for j := 0; j < count; j++ {
					addStr += string(slice[i-1])
				}
				if addStr != "" {
					res += addStr
				}
			}
		} else {
			res += string(compare)
		}
	}
	if escapeValue != "" {
		res += escapeValue
	}
	return res, nil
}
