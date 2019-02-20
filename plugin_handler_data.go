package sdk

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// PluginHandlerData defines the actions or triggers
type PluginHandlerData struct {
	RawName             string `yaml:"name"`
	Name                string `yaml:"-"` // This is the joined and camelled name for the action
	Title               string `yaml:"title"`
	Description         string `yaml:"description"`
	Input               ParamDataCollection
	Output              ParamDataCollection
	RequiresCustomTypes bool // Used to assist in generating correct imports
	Order               int  // Used to preserve the exact ordering in the spec
	InputSchema         JSONSchema
	OutputSchema        JSONSchema
	// Things that are only used to make parsing templates simpler
	PackageRoot string `yaml:"-"`
	HasInterval bool   `yaml:"-"`
}

// PluginHandlerCollection defines a type we can control marshal behavior with
type PluginHandlerCollection map[string]PluginHandlerData

// UnmarshalYAML lets us do some magic to set an order since go maps lack it
func (phc *PluginHandlerCollection) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	if phc == nil {
		phc = &PluginHandlerCollection{}
	}
	m := make(map[string]PluginHandlerData)
	if err := unmarshal(&m); err != nil {
		return fmt.Errorf("Unable to unmarshal yaml into PluginHandlerCollection: %s", err)
	}
	// Unmarshal again into an ordered list (in order to set Order property)
	op := make(yaml.MapSlice, 0)
	if err := unmarshal(&op); err != nil {
		return fmt.Errorf("Unable to unmarshal yaml into ordered parameters: %s", err)
	}
	for i, p := range op {
		key, ok := p.Key.(string)
		if !ok {
			return fmt.Errorf("Unable to unmarshal yaml parameters: Parameter name was not a string")
		}
		val, ok := m[key]
		if !ok {
			return fmt.Errorf("Could not unmarshal yaml parameter: %s key was somehow not in PluginHandlerCollection", key)
		}
		val.RawName = key
		val.Name = UpperCamelCase(key)
		val.Order = i + 1 // Set the order
		m[key] = val
	}
	*phc = m
	return nil
}
