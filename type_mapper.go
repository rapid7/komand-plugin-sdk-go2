package sdk

import (
	"log"
	"strings"
)

// TypeMapper handles renaming types from the spec when they don't line up with go types by name
type TypeMapper struct {
	CustomTypes    map[string]string
	StructureTypes map[string]string
}

// NewTypeMapper returns a new TypeMapper
func NewTypeMapper(s *PluginSpec) *TypeMapper { // TO FIGURE OUT object and bytes need some work, look at generate/conversion
	t := &TypeMapper{
		CustomTypes: map[string]string{
			"boolean":    "bool",
			"object":     "interface{}",
			"integer":    "int",
			"bytes":      "[]byte",
			"date":       "time.Time",
			"date-time":  "time.Time",
			"number":     "float64",
			"float":      "float64",
			"python":     "string",
			"password":   "string",
			"file":       "types.SDKFile",
			"credential": "credential",
		},
		StructureTypes: map[string]string{
			"token":          "types.CredentialToken",
			"passsword":      "types.CredentialPassword",
			"asymmetric_key": "types.CredentialAsymmetricKey",
		},
	}

	for name := range s.RawTypes {
		t.CustomTypes[name] = "types." + UpperCamelCase(name)
	}
	return t
}

// SpecTypeToGoType maps spec types to go types
func (t *TypeMapper) SpecTypeToGoType(pd ParamData) string {
	var isArr bool
	strippedType := pd.Type
	if strings.HasPrefix(pd.Type, "[]") {
		strippedType = pd.Type[2:]
		isArr = true
	}
	foundType, ok := t.CustomTypes[strippedType]
	if !ok {
		return pd.Type
	}
	if foundType == "credential" {
		if foundType, ok = t.StructureTypes[pd.Structure]; !ok {
			log.Printf("Cannot find Structure %s for type credential", pd.Structure)
			return pd.Type
		}
	}
	if isArr {
		foundType = "[]" + foundType
	}
	return foundType
}
