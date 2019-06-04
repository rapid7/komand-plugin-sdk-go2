package sdk

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// TypeMapper handles renaming types from the spec when they don't line up with go types by name
type TypeMapper struct {
	Types map[string]JSONSchema
}

// NewTypeMapper returns a new TypeMapper
func NewTypeMapper(s *PluginSpec) *TypeMapper {
	t := &TypeMapper{
		Types: map[string]JSONSchema{
			"bool":      Boolean,
			"boolean":   Boolean,
			"object":    Object,
			"int":       Integer,
			"integer":   Integer,
			"text":      String,
			"string":    String,
			"bytes":     Bytes,
			"date":      Date,
			"date-time": Date,
			"number":    Float,
			"float":     Float,
			"python":    Python,
			"password":  Password,
			"array":     Array,
			"file":      File,
			"credential_username_password": CredentialUsernamePassword,
			"credential_token":             CredentialToken,
			"credential_asymmetric_key":    CredentialAsymmetricKey,
			"credential_secret_key":        CredentialSecretKey,
		},
	}

	// Off the bat, copy all the top level types into our internal cache of knowledge
	// Once we have that list seeded, we can then enable the SDK to request type-by-type
	// a conversion from a "RawType" to a proper Type, which will include hydrating type
	// data and the json schemas
	for name := range s.RawTypes {
		customSchema := JSONSchema{
			ID:         name,
			Type:       "object",
			GoType:     "types." + UpperCamelCase(name),
			Properties: make(map[string]JSONSchema),
			BackRef:    true,
		}
		t.Types[name] = customSchema
	}
	return t
}

// ReferenceTypes returns a map of just the builtin composite types and any custom types
func (t *TypeMapper) ReferenceTypes() map[string]JSONSchema {
	m := make(map[string]JSONSchema)
	for k, v := range t.Types {
		if v.BackRef {
			m[k] = v
		}
	}
	return m
}

// RawTypeToType turns raw type data into a more refined sugar substance
func (t *TypeMapper) RawTypeToType(rawName string, rawData ParamDataCollection) (*TypeData, error) {
	td := &TypeData{
		Name:    UpperCamelCase(rawName),
		RawName: rawName,
	}
	if _, err := t.updateParams(rawData); err != nil {
		return nil, err
	}
	td.Fields = rawData
	td.SortedFields = sortParamData(rawData)
	return td, nil
}

// PopulateSchemasForTypeDatas pops all the schemas for the typedata
func (t *TypeMapper) PopulateSchemasForTypeDatas(typs map[string]TypeData) error {
	// First, build the schemas up
	for n, td := range typs {
		if err := t.PopulateSchemaForTypeData(&td); err != nil {
			return err
		}
		typs[n] = td
	}
	// now, backfill any definitions
	for n, td := range typs {
		td.Schema.AddDefinitions(t.ReferenceTypes())
		typs[n] = td
	}
	return nil
}

// PopulateSchemaForTypeData populates the schema for a single typedata
func (t *TypeMapper) PopulateSchemaForTypeData(td *TypeData) error {
	js := t.Types[td.RawName]
	js.Title = Title(td.RawName)
	if err := t.populateSchema(&js, td.Fields); err != nil {
		return err
	}
	t.Types[td.RawName] = js
	td.Schema = js
	return nil
}

// PopulateSchemaForPluginHandlerData populates the schemas for input and outputs of the plugin handler data
func (t *TypeMapper) PopulateSchemaForPluginHandlerData(phd *PluginHandlerData) error {
	is, err := t.GenerateSchemaForParamDataCollection(phd.Input)
	if err != nil {
		return nil
	}
	phd.InputSchema = *is
	os, err := t.GenerateSchemaForParamDataCollection(phd.Output)
	if err != nil {
		return nil
	}
	phd.OutputSchema = *os
	return nil
}

// GenerateSchemaForParamDataCollection generates a schema for a ParamDataCollection object
func (t *TypeMapper) GenerateSchemaForParamDataCollection(pdc ParamDataCollection) (*JSONSchema, error) {
	s := t.Types["object"] // Always based off type object
	s.Title = "Variables"  // Always named Variables
	s.Properties = make(map[string]JSONSchema)
	if err := t.populateSchema(&s, pdc); err != nil {
		return nil, err
	}
	s.AddDefinitions(t.ReferenceTypes())
	return &s, nil
}

// PopulateSchema is a method to fill in the json schema of a typedata
func (t *TypeMapper) populateSchema(js *JSONSchema, phd ParamDataCollection) error {
	// deep-map it
	for rawname, pd := range phd {
		typ := pd.RawType // we may need de-array this for certain lookups, so cache a copy of it
		isArray := false
		if strings.HasPrefix(typ, "[]") {
			// Prefixed by [] == it's an array type, we need to treat it + it's item type specially
			typ = typ[2:]
			isArray = true
		}
		// Make a copy of the schema, start filling it ut
		fschema := t.Types[typ]
		fschema.Title = pd.Title
		fschema.Required = nil
		// If this paramdata is required, track it in the parent-schemas list of requireds
		if pd.Required {

			if js.Required == nil {
				js.Required = []string{rawname}
			} else { // so many elses!
				js.Required = append(js.Required, rawname)
			}
			fmt.Println("---------")
			fmt.Println(pd.Title)
			fmt.Println("---------")
			fmt.Println(js.Required)
			fmt.Println("---------")
		}
		// If it needs a backref, fill it out and clear out things we don't use when using references
		// References should largely be empty, save a few helper props
		if fschema.BackRef {
			fschema.Ref = "#/definitions/" + fschema.ID
			// Get rid of the ID, this isn't the primary record for the type
			// we copied this schema when pulling it out of the map, so we don't
			// want 2 copies of the id floating around
			fschema.ID = ""
			fschema.Properties = nil
			fschema.Type = ""
		}
		// Set the array type items to be the same as the ID, it will eventually be turned into a ref
		if isArray {
			fschema.Type = "array"
			fschema.Properties = nil // arrays don't have properties, their items do
			// Need to see if this is a custom type ie a ref type
			if fschema.Ref != "" {
				// This was an array of custom types, not primitives
				fschema.Items = &JSONSchema{
					Ref: fschema.Ref,
				}
				fschema.Ref = "" // We copied the ref to the item, so clear out the arrays ref
			} else { // lol rare time where else is helpful
				// This was an array of primitves, not customs, so look up the right schema
				localItem := t.Types[typ]
				fschema.Items = &localItem // This one is easy - just use the primitive as the item type
			}
		}
		// If we haven't gotten a proper title on it by now, make our test effort default title
		if fschema.Title == "" {
			fschema.Title = Title(pd.RawName)
		}
		// Copy the remaining params
		fschema.Description = pd.Description
		fschema.Enum = pd.Enum
		fschema.Default = pd.Default
		fschema.Order = pd.Order
		// Store it on the props so it can be re-used with all this data later
		js.Properties[rawname] = fschema
	}
	return nil
}

func (t *TypeMapper) updateParams(rawData ParamDataCollection) (bool, error) {
	// Need to lookup the proper types for all the fields in the custom types object
	customTypes := false
	for name, param := range rawData {
		param.Type = t.SpecTypeToGoType(param.Type)
		param.EnumLiteral = make([]EnumData, len(param.Enum))
		if !customTypes && strings.Contains(param.Type, "types.") {
			customTypes = true
		}
		for i, e := range param.Enum { // TODO i think we can just type assert these to strings? but not sure if it'll always work...
			// So, here's how we're gonna do this: marshal the interface to json
			// this will give us a string representation of the value to write out as a literal
			// then we'll do const X = {{ Literal Value }}
			// The marshal approach might not be super efficient, but since generating a spec is a one-time
			// upfront cost, the convenience factor wins out for now.
			b, err := json.Marshal(e)
			if err != nil {
				return customTypes, err
			}
			param.EnumLiteral[i] = EnumData{
				Name:         param.Name + UpperCamelCase(string(b)),
				LiteralValue: string(b),
			}
		}
		rawData[name] = param // Godbless go for this feature
	}
	return customTypes, nil
}

// sortParamData takes the proprties for a generated class and sorts them based on if they are Embedded or not
// this is due to go wanting to see embeddables up front, then non embedded properties.
func sortParamData(pd ParamDataCollection) []ParamData {
	list := make([]ParamData, len(pd)) // <-
	var x int
	for _, v := range pd {
		list[x] = v
		x++
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Embed && !list[j].Embed // <-
	})
	return list
}

// SpecTypeToGoType maps spec types to go types
func (t *TypeMapper) SpecTypeToGoType(specType string) string {
	var isArr bool
	strippedType := specType
	if strings.HasPrefix(specType, "[]") {
		strippedType = specType[2:]
		isArr = true
	}
	foundType, ok := t.Types[strippedType]
	if !ok {
		// We don't have a remap for it, so use the original
		return specType
	}
	if isArr {
		return "[]" + foundType.GoType
	}
	return foundType.GoType
}
