package sdk

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rapid7/komand-plugin-sdk-go2/testspec"
	yaml "gopkg.in/yaml.v2"
)

func TestIVMSchemaGeneration(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(testspec.SpecIVMTrigger), s); err != nil {
		t.Fatal(err)
	}
	if err := PostProcessSpec(s); err != nil {
		t.Fatal(err)
	}
	if len(s.Types) == 0 {
		t.Fatal("parsing types failed")
	}
	if len(s.Triggers) == 0 {
		t.Fatal("parsing triggers failed")
	}
	js := s.Triggers["ivm_trigger"].OutputSchema
	b, _ := json.MarshalIndent(js, "    ", "")
	fmt.Println(string(b))
	t.Fatal("e")
}
