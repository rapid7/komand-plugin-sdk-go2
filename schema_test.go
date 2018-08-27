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
	for _, t := range s.Types {
		fmt.Println("------------------")
		fmt.Println(t.Name)
		//fmt.Println(t.RawName)
		b, _ := json.MarshalIndent(t.Schema, "", "    ")
		fmt.Println(string(b))
		//for fn, f := range t.Fields {
		//fmt.Println("\t" + fn)
		//fmt.Println("\t" + f.Name)
		//fmt.Println("\t" + f.RawName)

		//}
		fmt.Println("------------------")
		//break
	}
}
