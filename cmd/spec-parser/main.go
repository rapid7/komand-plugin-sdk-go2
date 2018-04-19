package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	sdk "github.com/rapid7/komand-plugin-sdk-go2"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Error, must provide path to spec file, including name of spec")
	}
	absPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	f, err := ioutil.ReadFile(path.Join(absPath))
	if err != nil {
		log.Fatal(err)
	}
	ps := &sdk.PluginSpec{}
	if err := yaml.Unmarshal(f, ps); err != nil {
		log.Fatalf("Error unmarshalling plugin spec: %s\n", err.Error())
	}
	if err := sdk.PostProcessSpec(ps); err != nil {
		log.Fatalf("Error processing plugin spec: %s\n", err.Error())
	}
}
