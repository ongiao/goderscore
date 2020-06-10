package goderscore

import (
	"html"
	"log"
	"regexp"
	"strings"
)

func CamelCase(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	stringSlice := strings.Split(strings.TrimSpace(reg.ReplaceAllString(str, " ")), " ")

	var res string

	for i, n := range stringSlice {
		if i == 0 {
			res += strings.ToLower(n)
		} else {
			res += strings.Title(n)
		}
	}

	return res
}

func Capitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}

func EndWith(str, target string) bool {
	return strings.HasSuffix(str, target)
}

func EndWithPosition(str, target string, position int) bool {
	if position < 0 || position >= len(str) {
		return false
	}

	return EndWith(str[:position], target)
}

func Escape(str string) string {
	return html.EscapeString(str)
}

func EscapeRegExp(str string) string {
	return regexp.QuoteMeta(str)
}

func KebabCase(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	processedStr := strings.TrimSpace(reg.ReplaceAllString(str, " "))

	var upperCaseIndexes []int

	for i, ch := range processedStr {
		if ch >= 65 && ch <= 90 {
			upperCaseIndexes = append(upperCaseIndexes, i)
		}
	}

	if len(upperCaseIndexes) > 0 {
		for _, index := range upperCaseIndexes {
			if index != 0 && processedStr[index-1] != ' ' {
				processedStr = processedStr[:index] + " " + processedStr[index:]
			}
		}
	}

	stringSlice := strings.Split(strings.ToLower(processedStr), " ")

	var res string

	for i, n := range stringSlice {
		if i == len(stringSlice) - 1 {
			res += n
		} else {
			res += n + "-"
		}
	}

	return res
}
