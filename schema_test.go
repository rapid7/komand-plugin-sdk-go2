package sdk

import (
	"encoding/json"
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestJSONSchemaGeneration(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(SpecBigFix), s); err != nil {
		t.Fatal(err)
	}
	if err := PostProcessSpec(s); err != nil {
		t.Fatal(err)
	}
	if len(s.Types) == 0 {
		t.Fatal("parsing actions failed")
	}
	b, _ := json.MarshalIndent(s.ConnectionSchema, "", "    ")
	fmt.Println(string(b))
	for _, t := range s.Actions {
		if t.InputSchema == nil || t.OutputSchema == nil {
			t.Fatal("action schemas were nil")
		}
	}
	for _, t := range s.Triggers {
		if t.InputSchema == nil || t.OutputSchema == nil {
			t.Fatal("trigger schemas were nil")
		}
	}
	if s.ConnectionSchema == nil {
		t.Fatal("connection schema is nil")
	}
}
