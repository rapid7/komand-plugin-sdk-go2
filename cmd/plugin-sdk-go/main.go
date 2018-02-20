package main

import (
	"flag"
	"log"

	sdk "github.com/rapid7/komand-plugin-sdk-go2"
)

func main() {
	specLoc := flag.String("spec", "", "The path to the spec file")
	packageRoot := flag.String("package", "", "The go package root for this plugin. Ex: github.com/<company_name>/plugins/<plugin_name>")
	flag.Parse()
	if *specLoc == "" {
		log.Fatal("Error, must provide path to spec including name")
	}
	if *packageRoot == "" {
		log.Fatal("Error, must provide a package root for the resulting go package. Ex: github.com/<company_name>/plugins/<plugin_name>")
	}
	g, err := sdk.NewGenerator(*specLoc, *packageRoot, false)
	if err != nil {
		log.Fatalf("error creating generator: %s", err)
	}
	// Now, MAKE IT HAPPEN
	if err := g.GeneratePlugin(); err != nil {
		log.Fatal(err)
	}
}
