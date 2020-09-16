package stringutil

import (
	"strings"
	"unicode/utf8"
)

const Empty = ""

func Equals(s1, s2 string) bool {
	return s1 == s2
}

func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

func NotEquals(s1, s2 string) bool {
	return !Equals(s1, s2)
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == Empty
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsEmpty(s string) bool {
	return s == Empty
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func Length(s string) int {
	return utf8.RuneCountInString(s)
}

func Contains(str, searchStr string) bool {
	return strings.Contains(str, searchStr)
}

func StartsWith(str, prefix string) bool {
	return startsWithIgnoreCase(str, prefix, false)
}

func StartsWithIgnoreCase(str, prefix string) bool {
	return startsWithIgnoreCase(str, prefix, true)
}

func startsWithIgnoreCase(str, prefix string, ignore bool) bool {
	if len(prefix) > len(str) {
		return false
	}
	s := str[:len(prefix)]
	if ignore {
		return strings.EqualFold(s, prefix)
	}
	return s == prefix
}
