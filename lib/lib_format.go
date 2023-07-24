package lib

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// special chars replaced with _
func SafeName(value string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(value, "_")
}

func PrivVarName(value string) string {
	return strings.ToLower(SafeName(value))
}

func PubVarName(value string) string {
	caser := cases.Title(language.BritishEnglish)
	a := caser.String(SafeName(value))
	return a
}

func FuncName(value string) string {
	return strcase.ToCamel(strings.ReplaceAll(SafeName(value), "_", " "))
}

func DisplayName(value string) string {
	caser := cases.Title(language.BritishEnglish)
	return caser.String(strings.ReplaceAll(value, "_", " "))
}
