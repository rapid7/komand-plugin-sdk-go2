package sdk

import (
	"strings"
)

// JSONSchema describes a JSON schema.
// It's used for marshalling types to and from JSON Schemas.
type JSONSchema struct {
	ID          string                `json:"id,omitempty" yaml:"-"`
	Ref         string                `json:"$ref,omitempty"`
	Type        string                `json:"type,omitempty"`
	Title       string                `json:"title,omitempty"`
	DisplayType string                `json:"displayType,omitempty"` // DisplayType will override Type when showing a Badge for things like Bytes and Date, due to the way those work
	Description string                `json:"description,omitempty"`
	Items       *JSONSchema           `json:"items,omitempty"`
	Properties  map[string]JSONSchema `json:"properties,omitempty"`
	Required    []string              `json:"required,omitempty"`
	Default     interface{}           `json:"default,omitempty"`
	Enum        []interface{}         `json:"enum,omitempty"`
	Definitions map[string]JSONSchema `json:"definitions,omitempty"`
	Format      string                `json:"format,omitempty"`
	Order       int                   `json:"order,omitempty"`
	OneOf       []JSONSchema          `json:"oneOf,omitempty"`
	GoType      string                `json:"-"` // only used internally for generation of plugins - python doesn't use this
	BackRef     bool                  `json:"-"` // only used internally
}

// RefTypeName returns the type name
func (s JSONSchema) RefTypeName() string {
	// The ref type name of arrays is their items type, not their type
	if s.Type == "array" && s.Items != nil {
		if strings.HasPrefix(s.Items.Ref, "#/definitions/") {
			return strings.TrimPrefix(s.Items.Ref, "#/definitions/")
		}
		return ""
	}
	// the ref type name of normal objects or types is their ref
	if strings.HasPrefix(s.Ref, "#/definitions/") {
		return strings.TrimPrefix(s.Ref, "#/definitions/")
	}
	return ""
}

// IsObjectType returns true if an object type
func (s JSONSchema) IsObjectType() bool {
	return s.Type == "object"
}

// IsArrayType is true if array
func (s JSONSchema) IsArrayType() bool {
	return s.Type == "array"
}

// Equal compares two schemas, and returns true if they are equivalent
// copied from old world komand
func (s JSONSchema) Equal(other JSONSchema) bool {
	if s.Type != other.Type || s.Format != other.Format {
		return false
	}
	if s.BackRef {
		if other.BackRef && other.Ref == s.Ref {
			return true
		}
		return false
	}
	if s.IsArrayType() {
		if !other.IsArrayType() {
			return false
		}
		if s.Items != other.Items {
			if s.Items != nil && other.Items != nil {
				return s.Items.Equal(*other.Items)
			}
			return false
		}
		return true
	}
	if s.IsObjectType() {
		if !other.IsObjectType() {
			return false
		}
		if s.Properties == nil && other.Properties == nil {
			return true
		}
		return compareMaps(s.Properties, other.Properties)
	}
	return true
}

// copied from old world komand
func compareMaps(source map[string]JSONSchema, dest map[string]JSONSchema) bool {
	if source == nil && dest != nil {
		return false
	}
	if source != nil && dest == nil {
		return false
	}
	if len(source) != len(dest) {
		return false
	}
	for key, val := range source {
		destVal, ok := dest[key]
		if !ok {
			return false
		}
		if !val.Equal(destVal) {
			return false
		}
	}
	return true
}

// IsSubset compares two schemas, and returns true if s is equivalent or a subset of superset
// copied from old world komand
func (s JSONSchema) IsSubset(superset JSONSchema) bool {
	if s.Type != superset.Type || s.Format != superset.Format {
		return false
	}
	if s.BackRef {
		if superset.BackRef && superset.Ref == s.Ref {
			return true
		}
		return false
	}
	if s.IsArrayType() {
		if !superset.IsArrayType() {
			return false
		}
		if s.Items != superset.Items {
			if s.Items != nil && superset.Items != nil {
				return s.Items.IsSubset(*superset.Items)
			}
			return false
		}
		return true
	}
	if s.IsObjectType() {
		if !superset.IsObjectType() {
			return false
		}
		if s.Properties == nil && superset.Properties == nil {
			return true
		}
		return isMapSubset(s.Properties, superset.Properties)
	}
	return true
}

// copied from old world komand
func isMapSubset(source map[string]JSONSchema, superset map[string]JSONSchema) bool {
	if source == nil && superset != nil {
		return true
	}
	if source != nil && superset == nil {
		return false
	}
	if len(source) > len(superset) {
		return false
	}
	for key, val := range source {
		supersetVal, ok := superset[key]
		if !ok {
			return false
		}
		if !val.IsSubset(supersetVal) {
			return false
		}
	}
	return true
}

// AddDefinitions adds definitions only *if* the schema needs them
// copied from old world komand
func (s *JSONSchema) AddDefinitions(defs map[string]JSONSchema) {
	if s.Definitions == nil {
		s.Definitions = make(map[string]JSONSchema)
	}
	for _, name := range DetectRefs(s, defs) {
		def, ok := defs[name]
		if ok {
			// We don't set the id in the definition
			// it's only used while processing to build backrefs where needed
			def.ID = ""
			s.Definitions[name] = def
		}
	}
}

// DetectRefs detect json schema references
// copied from old world komand
func DetectRefs(s *JSONSchema, defs map[string]JSONSchema) []string {
	var refs []string
	if s.BackRef {
		ref := s.RefTypeName()
		refs = append(refs, ref)
		def, ok := defs[ref]
		if ok {
			refs = append(refs, DetectRefs(&def, defs)...)
		}
	}
	if s.Items != nil {
		refs = append(refs, DetectRefs(s.Items, defs)...)
	}
	if s.Properties != nil {
		for _, p := range (*s).Properties {
			refs = append(refs, DetectRefs(&p, defs)...)
		}
	}
	return refs
}
