package sdk

import (
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/rapid7/komand-plugin-sdk-go2/testspec"
)

func TestJSONSchemaGeneration(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(testspec.SpecBigFix), s); err != nil {
		t.Fatal(err)
	}
	if err := PostProcessSpec(s); err != nil {
		t.Fatal(err)
	}
	if len(s.Types) == 0 {
		t.Fatal("parsing actions failed")
	}
	for _, ac := range s.Actions {
		if ac.InputSchema.Equal(JSONSchema{}) || ac.OutputSchema.Equal(JSONSchema{}) {
			t.Fatal("action schemas were empty")
		}
	}
	for _, tr := range s.Triggers {
		if tr.InputSchema.Equal(JSONSchema{}) || tr.OutputSchema.Equal(JSONSchema{}) {
			t.Fatal("trigger schemas were nil")
		}
	}
	if s.ConnectionSchema.Equal(JSONSchema{}) {
		t.Fatal("connection schema is nil")
	}
}
