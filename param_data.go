package sdk

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// ParamData are the details that make up each input/output/trigger param
// Not all data is always filled in, it's context sensitive, which is unfortunate but
// I'm willing to accept it for this since everything else is greatly simplified
type ParamData struct {
	RawName     string        `yaml:"name"`
	Name        string        `yaml:"-"` // This is the joined and camelled name for the param
	RawType     string        `yaml:"-"` // This is the original type as it was in the yaml
	Type        string        `yaml:"type"`
	Title       string        `yaml:"title"`
	Required    bool          `yaml:"required"`
	Description string        `yaml:"description"`
	Enum        []interface{} `yaml:"enum"`
	Default     interface{}   `yaml:"default"`
	Embed       bool          `yaml:"embed"`
	Nullable    bool          `yaml:"nullable"`
	Order       int
	Schema      JSONSchema

	// Things that are used for background help
	EnumLiteral []EnumData `yaml:"-"`
}

// ParamDataCollection defines a type we can control marshal behavior with
type ParamDataCollection map[string]ParamData

// UnmarshalYAML lets us do some magic to set an order since go maps lack it
func (pdc *ParamDataCollection) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	if pdc == nil {
		pdc = &ParamDataCollection{}
	}
	m := make(map[string]ParamData)
	if err := unmarshal(&m); err != nil {
		return fmt.Errorf("Unable to unmarshal yaml into ParamDataCollection: %s", err)
	}
	// Unmarshal again into an ordered list (in order to set Order property)
	orderedParameters := make(yaml.MapSlice, 0)
	if err := unmarshal(&orderedParameters); err != nil {
		return fmt.Errorf("Unable to unmarshal yaml into ordered parameters: %s", err)
	}
	for i, p := range orderedParameters {
		key, ok := p.Key.(string)
		if !ok {
			return fmt.Errorf("Unable to unmarshal yaml parameters: Parameter name was not a string")
		}
		val, ok := m[key]
		if !ok {
			return fmt.Errorf("Could not unmarshal yaml parameter: %s key was somehow not in ParamDataCollection", key)
		}
		val.RawType = val.Type
		val.RawName = key
		val.Name = UpperCamelCase(key)
		val.Order = i + 1 // Set the order
		m[key] = val
	}
	*pdc = m
	return nil
}

// EnumData is used to parse and write out enums
type EnumData struct {
	Name         string `yaml:"-"`
	LiteralValue string `yaml:"-"`
}

// TypesInternalType is a special hack for the types package. Because in all other places we need to reference X via types.X
// we prefix it ahead of time. But types that use or refer to other types in the types package don't need it.
// This method, used only from the code generators for types, is to make sure types don't use the pacakge name internally in the package
func (p ParamData) TypesInternalType() string {
	if i := strings.Index(p.Type, "types."); i > -1 {
		return strings.Replace(p.Type, "types.", "", -1)
	}
	return p.Type
}
