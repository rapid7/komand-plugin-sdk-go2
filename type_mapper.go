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
			"credential_username_password": CredentialAsymmetricKey,
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
		fmt.Printf("^^^^^^^^^^^^^^^^^^^^ %s\n", name)
		t.Types[name] = customSchema
	}
	/*for name, typ := range s.RawTypes {
		// get the existing record
	}*/
	return t
}

// RawTypeToType is
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

// PopulateSchemas pops all the schemas
func (t *TypeMapper) PopulateSchemas(typs map[string]TypeData) error {
	for n, td := range typs {
		if err := t.PopulateSchema(&td); err != nil {
			return err
		}
		typs[n] = td
	}
	return nil
}

// PopulateSchema is
func (t *TypeMapper) PopulateSchema(td *TypeData) error {
	js := t.Types[td.RawName]
	js.Title = Title(td.RawName)
	// Now, deep-map it
	for rawname, pd := range td.Fields {
		typ := pd.RawType
		isArray := false
		if strings.HasPrefix(typ, "[]") {
			typ = typ[2:]
			isArray = true
		}
		fschema := t.Types[typ]
		fmt.Println(fschema)
		// fill it out
		if fschema.BackRef {
			fschema.Ref = "#/definitions/" + fschema.ID
			// Get rid of the ID, this isn't the primary record for the type
			// we copied this schema when pulling it out of the map, so we don't
			// want 2 copies of the id floating around
			fschema.ID = ""
		}
		// Set the array type items to be the same as the ID, it will eventually be turned into a ref
		if isArray {
			fschema.Type = "array"
			fschema.Properties = nil // arrays don't have properties
			// Need to see if this is a custom type ie a ref type
			if fschema.Ref != "" {
				// This was an array of custom types, not primitives
				fschema.Items = &JSONSchema{
					Ref: fschema.Ref,
				}
				fschema.Ref = ""
			} else { // lol rare time where else is helpful
				// This was an array of primitves, not customs, so look up the right schema
				localItem := t.Types[typ]
				fschema.Items = &localItem
			}
		}
		fschema.Title = Title(pd.RawName)
		fschema.Description = pd.Description
		fschema.Enum = pd.Enum
		fschema.Default = pd.Default
		fschema.Order = pd.Order
		// Store it on the props
		js.Properties[rawname] = fschema
	}
	t.Types[td.RawName] = js
	td.Schema = js
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
		for i, e := range param.Enum {
			//fmt.Printf("--------------------- %T -----------------\n", e) // TODO am i just doing dumbshit here?
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

// Boolean is a boolean
var Boolean = JSONSchema{
	Type:   "boolean",
	GoType: "bool",
}

// Object is an object
var Object = JSONSchema{
	Type:   "object",
	GoType: "interface{}",
}

// Integer are integers
var Integer = JSONSchema{
	Type:   "integer",
	GoType: "int",
}

// Bytes are byte strings
var Bytes = JSONSchema{
	Type:        "string",
	GoType:      "[]byte",
	Format:      "bytes",
	DisplayType: "bytes",
}

// Date is a date
var Date = JSONSchema{
	Type:        "string",
	GoType:      "time.Time",
	Format:      "date-time",
	DisplayType: "date",
}

// Float is a floating point number
var Float = JSONSchema{
	Type:   "number",
	GoType: "float64",
}

// Python is python code
var Python = JSONSchema{
	Type:        "string",
	GoType:      "string",
	Format:      "python",
	DisplayType: "python",
}

// Password is a password type
var Password = JSONSchema{
	Type:        "string",
	GoType:      "string",
	Format:      "password",
	DisplayType: "password",
}

// String is a string
var String = JSONSchema{
	Type:   "string",
	GoType: "string",
}

// Array is generic array
var Array = JSONSchema{
	Type:   "array",
	GoType: "[]interface{}",
}

// File is a file
var File = JSONSchema{
	ID:          "file",
	Title:       "File",
	Description: "File Object",
	Type:        "object",
	GoType:      "types.SDKFile",
	BackRef:     true,
	Ref:         "#/definitions/file",
	Properties: map[string]JSONSchema{
		"filename": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Filename",
			Description: "Name of file",
		},
		"content": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Format:      "bytes",
			Title:       "Content",
			Description: "File contents",
		},
	},
}

// CredentialUsernamePassword is a CC credential type for managing login names and passwords.
var CredentialUsernamePassword = JSONSchema{
	ID:          "credential_username_password",
	Title:       "Credential: Username and Password",
	Description: "A username and password combination",
	Type:        "object",
	GoType:      "types.CredentialUsernamePassword",
	Required:    []string{"username", "password"},
	BackRef:     true,
	Ref:         "#/definitions/credential_username_password",
	Properties: map[string]JSONSchema{
		"username": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Username",
			Description: "The username to log in with",
		},
		"password": JSONSchema{
			Title:       "Password",
			Description: "The password",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}

// CredentialAsymmetricKey is a CC credential type for managing access tokens
var CredentialAsymmetricKey = JSONSchema{
	ID:          "credential_asymmetric_key",
	Title:       "Credential: Asymmetric Key",
	Description: "A shared key",
	Type:        "object",
	GoType:      "types.CredentialAsymmetricKey",
	Required:    []string{"privateKey"},
	BackRef:     true,
	Ref:         "#/definitions/credential_asymmetric_key",
	Properties: map[string]JSONSchema{
		"privateKey": JSONSchema{
			Title:       "Private Key",
			Description: "The private key",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}

// CredentialToken is a CC credential type for managing a token and an an optional domain
var CredentialToken = JSONSchema{
	ID:          "credential_token",
	Title:       "Credential: Token",
	Description: "A pair of a token, and an optional domain",
	Type:        "object",
	GoType:      "types.CredentialToken",
	Required:    []string{"token"},
	BackRef:     true,
	Ref:         "#/definitions/credential_token",
	Properties: map[string]JSONSchema{
		"token": JSONSchema{
			Title:       "Token",
			Description: "The shared token",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
		"domain": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Domain",
			Description: "The domain for the token",
		},
	},
}

// CredentialSecretKey is a CC credential type for managing access tokens
var CredentialSecretKey = JSONSchema{
	ID:          "credential_secret_key",
	Title:       "Credential: Secret Key",
	Description: "A shared secret key",
	Type:        "object",
	GoType:      "types.CredentialSecretKey",
	Required:    []string{"secretKey"},
	BackRef:     true,
	Ref:         "#/definitions/credential_secret_key",
	Properties: map[string]JSONSchema{
		"secretKey": JSONSchema{
			Title:       "Secret Key",
			Description: "The shared secret key",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}
