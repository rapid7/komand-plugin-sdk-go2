package sdk

import (
	"testing"

	"github.com/rapid7/komand-plugin-sdk-go2/testspec"
	yaml "gopkg.in/yaml.v2"
)

func TestIDRSchemaGeneration(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(testspec.SpecIDRTrigger), s); err != nil {
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
}
