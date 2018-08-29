package sdk

import (
	"encoding/json"
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestIVMSchemaGeneration(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(SpecExchange), s); err != nil {
		t.Fatal(err)
	}
	if err := PostProcessSpec(s); err != nil {
		t.Fatal(err)
	}
	if len(s.Types) == 0 {
		t.Fatal("parsing actions failed")
	}
	/*for _, ty := range s.Types {
		b, _ := json.MarshalIndent(ty.Schema, "", "    ")
		fmt.Println(string(b))
		fmt.Println("---------------------")
	}*/
	for _, tr := range s.Triggers {
		fmt.Printf("--------------- %s --------------\n", tr.Name)
		b, _ := json.MarshalIndent(tr.OutputSchema, "", "    ")
		fmt.Println(string(b))
		fmt.Println("---------------------")
	}
}
