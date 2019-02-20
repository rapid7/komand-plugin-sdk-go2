package sdk

import (
	"bytes"
	"regexp"
	"strings"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

//UpperCamelCase turns strings -> upper camel case
func UpperCamelCase(src string) string {
	return CamelCase(src, true)
}

// CamelCase camel cases things
func CamelCase(src string, startWithUpper bool) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 || startWithUpper {
			if ok := commonInitialisms[strings.ToUpper(string(val))]; ok {
				// Instead of Titling it, make it uppercased all the way
				chunks[idx] = bytes.ToUpper(val)
				continue
			}
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}

// CamelCase camel cases things
func Title(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, []byte(" ")))
}

// Actually taken from: https://github.com/serenize/snaker/blob/master/snaker.go
// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}
