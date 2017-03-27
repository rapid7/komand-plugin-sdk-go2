package sdk

import (
	"bytes"
	"encoding/json"
	"regexp"
	"sort"
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

// This is a light weight helper to handle some post-processing boilerplate after parsing a spec
// We need to convert the spec param names into go friendly names, and lookup the proper type
func updateParams(data map[string]ParamData, t *TypeMapper) error {
	for name, param := range data {
		param.RawName = name
		param.Name = UpperCamelCase(name)
		param.Type = t.SpecTypeToGoType(param.Type)
		param.EnumLiteral = make([]EnumData, len(param.Enum))
		for i, e := range param.Enum {
			// So, here's how we're gonna do this: marshal the interface to json
			// this will give us a string representation of the value to write out as a literal
			// then we'll do const X = {{ Literal Value }}
			b, err := json.Marshal(e)
			if err != nil {
				return err
			}
			param.EnumLiteral[i] = EnumData{
				Name:         param.Name + UpperCamelCase(string(b)),
				LiteralValue: string(b),
			}
		}
		data[name] = param // Godbless go for this feature
	}
	return nil
}

// Hello
// I hope that this method in no way becomes the unadulterated shitshow that it probably will
// Sincerely
// Everyone who ever tried to do something useful in the spur of the moment, ever
func sortParamData(pd map[string]ParamData) []ParamData {
	list := make([]ParamData, len(pd)) // <-
	var i int
	for _, v := range pd {
		list[i] = v
		i++
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Embed && !list[j].Embed // <-
	})
	return list
}
