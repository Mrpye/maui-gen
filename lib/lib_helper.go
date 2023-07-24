package lib

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func CheckValidUrl(host string) error {
	if strings.Contains(strings.ToLower(host), "http://") || strings.Contains(strings.ToLower(host), "https://") {
		return nil
	} else {
		return fmt.Errorf("host must be a valid URL")
	}
}
func Minus(a, b int) int {
	return a - b
}
func Plus(a, b int) int {
	return a + b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	return a / b
}

func Concat(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//strip the schema from url
func GetDomainOrIP(url_str string) string {
	u, _ := url.Parse(url_str)
	host, _, _ := net.SplitHostPort(u.Host)
	return host
}
func GetPortString(url_str string) string {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	return port
}
func GetPortInt(url_str string) int {
	u, _ := url.Parse(url_str)
	_, port, _ := net.SplitHostPort(u.Host)
	intVar, _ := strconv.Atoi(port)
	return intVar
}
func Clean(str string, replace string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err == nil {
		return re.ReplaceAllString(str, replace)
	}
	return str
}

func RemoveURL(str string) string {
	str = strings.ReplaceAll(str, "https://", "")
	str = strings.ReplaceAll(str, "http://", "")
	return str
}

func InsertString(orig []string, index int, value string) []string {
	if index >= len(orig) {
		return append(orig, value)
	}
	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value

	return orig
}

func CastStringToType(input_type string, value interface{}) interface{} {
	if input_type == "bool" {
		b1, _ := strconv.ParseBool(value.(string))
		return b1
	} else if input_type == "int" {
		b1, _ := strconv.ParseInt(value.(string), 10, 64)
		return b1
	} else if input_type == "float" {
		b1, _ := strconv.ParseFloat(value.(string), 10)
		return b1
	}
	return value
}
