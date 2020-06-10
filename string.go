package goderscore

import (
	"html"
	"log"
	"regexp"
	"strconv"
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

	allUpperCase := true
	var upperCaseIndexes []int
	for i, ch := range processedStr {
		if ch >= 65 && ch <= 90 {
			upperCaseIndexes = append(upperCaseIndexes, i)
		} else if ch == 32 {
		} else {
			allUpperCase = false
		}
	}

	if allUpperCase {
		return strings.ReplaceAll(strings.ToLower(processedStr), " ", "-")
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

func LowerCase(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	processedStr := strings.TrimSpace(reg.ReplaceAllString(str, " "))

	allUpperCase := true
	var upperCaseIndexes []int
	for i, ch := range processedStr {
		if ch >= 65 && ch <= 90 {
			upperCaseIndexes = append(upperCaseIndexes, i)
		} else if ch == 32 {
		} else {
			allUpperCase = false
		}
	}

	if allUpperCase {
		return strings.ToLower(processedStr)
	}

	if len(upperCaseIndexes) > 0 {
		for _, index := range upperCaseIndexes {
			if index != 0 && processedStr[index-1] != ' ' {
				processedStr = processedStr[:index] + " " + processedStr[index:]
			}
		}
	}

	return strings.ToLower(processedStr)
}

func LowerFirst(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func Pad(str string, length int, chars string) string {
	if length <= len(str) {
		return str
	}

	repeatCharsHead := strings.Repeat(chars, length)
	repeatCharsTail := strings.Repeat(chars, length)
	startIndex := (length - len(str)) / 2

	return repeatCharsHead[:startIndex] + str + repeatCharsTail[:length - startIndex - len(str)]
}

func PadEnd(str string, length int, chars string) string {
	if length <= len(str) {
		return str
	}

	repeatChars := strings.Repeat(chars, length)

	return str + repeatChars[:length - len(str)]
}

func PadStart(str string, length int, chars string) string {
	if length <= len(str) {
		return str
	}

	repeatChars := strings.Repeat(chars, length)

	return repeatChars[:length - len(str)] + str
}

func ParseInt(str string, radix int) int {
	bitSize := 32 << (^uint(0) >> 63)
	i, err := strconv.ParseInt(str, radix, bitSize)

	if err != nil {
		log.Fatal("Parse int error!")
	}

	return int(i)
}

func Repeat(str string, n int) string {
	return strings.Repeat(str, n)
}

func Replace(str string, old, replacement string, n int) string {
	return strings.Replace(str, old, replacement, n)
}

func SnakeCase(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	processedStr := strings.TrimSpace(reg.ReplaceAllString(str, " "))

	allUpperCase := true
	var upperCaseIndexes []int
	for i, ch := range processedStr {
		if ch >= 65 && ch <= 90 {
			upperCaseIndexes = append(upperCaseIndexes, i)
		} else if ch == 32 {
		} else {
			allUpperCase = false
		}
	}

	if allUpperCase {
		return strings.ReplaceAll(strings.ToLower(processedStr), " ", "_")
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
			res += n + "_"
		}
	}

	return res
}

func Split(str, separator string) []string {
	return strings.Split(str, separator)
}

func SplitLimit(str, separator string, limit int) []string {
	res := strings.Split(str, separator)

	return res[:limit]
}

func StartCase(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	processedStr := strings.TrimSpace(reg.ReplaceAllString(str, " "))

	allUpperCase := true
	var upperCaseIndexes []int
	for i, ch := range processedStr {
		if ch >= 65 && ch <= 90 {
			upperCaseIndexes = append(upperCaseIndexes, i)
		} else if ch == 32 {
		} else {
			allUpperCase = false
		}
	}

	if allUpperCase {
		return processedStr
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
			res += strings.Title(n)
		} else {
			res += strings.Title(n) + " "
		}
	}

	return res
}

func StartsWith(str, target string, position int) bool {
	if position >= len(str) {
		return false
	}

	if position < 0 {
		position = 0
	}

	return strings.Index(str[position:], target) == 0
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func Trim(str, chars string) string {
	return strings.Trim(str, chars)
}

func TrimEnd(str, chars string) string {
	return strings.TrimRight(str, chars)
}

func TrimStart(str, chars string) string {
	return strings.TrimLeft(str, chars)
}