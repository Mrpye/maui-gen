package lib

import "strings"

func NOT(a bool) bool {
	return !a
}
func OR(a bool, b bool) bool {
	return a || b
}
func AND(a bool, b bool) bool {
	return a && b
}

func StringContainsStringListItem(string_list string, str string) bool {
	arr := strings.Split(string_list, ",")
	for _, a := range arr {
		if strings.Contains(str, a) {
			return true
		}
	}
	return false
}
