package sdk

import "strings"

// PostProcessSpec does some minor post-processing on the spec object to fill a few things in that make
// template generation easier
func PostProcessSpec(s *PluginSpec) error {
	// Create a new TypeManager and feed it some metadata about the spec, for it's own benefits
	t := NewTypeMapper(s)
	s.TypeMapper = t

	// Handle any custom types
	// We'll both populate Types AND update RawTypes so the original source is correct w/r/t the downstream source
	s.Types = make(map[string]TypeData)
	for name, data := range s.RawTypes {
		td, err := t.RawTypeToType(name, data)
		if err != nil {
			return err
		}
		s.RawTypes[name] = data
		s.Types[name] = *td
	}
	// now fill in all the schemas
	if err := t.PopulateSchemasForTypeDatas(s.Types); err != nil {
		return err
	}

	// fill in the connection names
	// Do this one out long form since we need the special case of building the data key
	for name, param := range s.Connection {
		param.RawName = name
		param.Name = UpperCamelCase(name)
		specType := param.Type
		param.Type = t.SpecTypeToGoType(param.Type)
		param.RawType = specType
		// Make sure to import the plugins own types package when generating the connection
		if !s.ConnectionRequiresCustomTypes && strings.Contains(param.Type, "types.") {
			s.ConnectionRequiresCustomTypes = true
		}
		s.Connection[name] = param
		// This is all stupid hacky but we need some way to hash the params and have a consistent key
		// for the connection in the hash for looking up via the data incoming with the message
		if param.Type == "string" && specType != "python" && specType != "password" {
			if s.ConnectionDataKey != "" {
				s.ConnectionDataKey += " + "
			}
			s.ConnectionDataKey += "c." + param.Name
		}
	}
	if s.ConnectionDataKey == "" {
		// Default it to the literal value of an empty string for now
		// This could be because there were no string params - an issue to solve
		// Or because it doesn't use a connection, which is totally fine
		// TODO if there is no connection to generate, skip the whole connection pkg?
		s.ConnectionDataKey = `""`
	}
	// Now for Connection Schema
	cs, err := t.GenerateSchemaForParamDataCollection(s.Connection)
	if err != nil {
		return err
	}
	s.ConnectionSchema = *cs
	// fill in the action and trigger data
	for name, action := range s.Actions {
		if action.RawName != "" && action.Title == "" {
			// Very old plugins used name, not title
			// We use RawName for different purposes in the gosdk
			// swap the values around if no title, but name was set
			action.Title = action.RawName
		}
		//action.RawName = name // not set in the yaml this way, but set for the benefit of the template
		action.Name = UpperCamelCase(name)
		action.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		customTypesInput, err := t.updateParams(action.Input)
		if err != nil {
			return err
		}
		customTypesOutput, err := t.updateParams(action.Output)
		if err != nil {
			return err
		}
		if customTypesInput || customTypesOutput {
			action.RequiresCustomTypes = true
		}
		// add schema to actions
		if err := t.PopulateSchemaForPluginHandlerData(&action); err != nil {
			return err
		}
		s.Actions[name] = action
	}

	for name, trigger := range s.Triggers {
		if trigger.RawName != "" && trigger.Title == "" {
			// Very old plugins used name, not title
			// We use RawName for different purposes in the gosdk
			// swap the values around if no title, but name was set
			trigger.Title = trigger.RawName
		}
		//trigger.RawName = name // not set in the yaml this way, but set for the benefit of the template
		trigger.Name = UpperCamelCase(name)
		trigger.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		customTypesInput, err := t.updateParams(trigger.Input)
		if err != nil {
			return err
		}
		customTypesOutput, err := t.updateParams(trigger.Output)
		if err != nil {
			return err
		}
		if customTypesInput || customTypesOutput {
			trigger.RequiresCustomTypes = true
		}
		if _, ok := trigger.Input["interval"]; ok {
			trigger.HasInterval = true
		}
		// add schema to actions
		if err := t.PopulateSchemaForPluginHandlerData(&trigger); err != nil {
			return err
		}
		s.Triggers[name] = trigger
	}
	if s.HTTP.Port == 0 {
		s.HTTP.Port = 10001
	}
	if s.HTTP.ReadTimeout == 0 {
		s.HTTP.ReadTimeout = 60 * 10 // Default read timeout is 10 mins
	}
	if s.HTTP.WriteTimeout == 0 {
		s.HTTP.WriteTimeout = 60 * 10 // default write timeout is 10 mins
	}
	return nil
}
