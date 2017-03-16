package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
)

// PluginSpec is
type PluginSpec struct {
	PluginSpecVersion string                          `yaml:"plugin_spec_version"`
	Name              string                          `yaml:"name"`
	Title             string                          `yaml:"title"`
	Description       string                          `yaml:"description"`
	Version           string                          `yaml:"version"`
	Vendor            string                          `yaml:"vendor"`
	Tags              []string                        `yaml:"tags"`
	Icon              string                          `yaml:"icon"`
	Help              string                          `yaml:"help"`
	Connection        map[string]ParamData            `yaml:"connection"`
	RawTypes          map[string]map[string]ParamData `yaml:"types"`
	Types             map[string]TypeData             `yaml:"-"`
	Triggers          map[string]PluginHandlerData    `yaml:"triggers"`
	Actions           map[string]PluginHandlerData    `yaml:"actions"`

	// Things that are not part of the spec, but still important
	PackageRoot string `yaml:"package_root"`
	HTTP        HTTP   `yaml:"http"`

	// Things that are for background help
	TypeMapper   *TypeMapper           `yaml:"-"`
	Enums        map[string][]EnumData `yaml:"-"`
	SpecLocation string                `yaml:"-"`
}

// ParamData are the details that make up each input/output/trigger param
// Not all data is always filled in, it's context sensitive, which is unfortunate but
// I'm willing to accept it for this since everything else is greatly simplified
type ParamData struct {
	RawName     string        `yaml:"name"`
	Name        string        `yaml:"-"`    // This is the joined and camelled name for the param
	Type        string        `yaml:"type"` // This needs some work - we have to convert some types to go types, like object/[]object...
	Required    bool          `yaml:"required"`
	Description string        `yaml:"description"`
	Enum        []interface{} `yaml:"enum"`
	Default     interface{}   `yaml:"default"`
}

// PluginHandlerData defines the actions or triggers
type PluginHandlerData struct {
	RawName     string `yaml:"name"`
	Name        string `yaml:"-"` // This is the joined and camelled name for the action
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Input       map[string]ParamData
	Output      map[string]ParamData

	// Things that are only used to make parsing templates simpler
	PackageRoot string `yaml:"-"`
	// TODO switch to a list of bonus import statements or see later TODO about using goimports
	NeedsTypes bool `yaml:"-"`
	NeedsTime  bool `yaml:"-"`
}

// TypeData defines the custom types. Much of the data is pulled via a parent-key, so we don't parse much from yaml at all.
// Instead, we post-process populate it for the benefit of the template
type TypeData struct {
	RawName string               `yaml:"-"`
	Name    string               `yaml:"-"`
	Fields  map[string]ParamData `yaml:"-"`
}

// HTTP Defines the settings for the plugins http server
type HTTP struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
}

// EnumData defines enumerated values
type EnumData struct {
	Name  string      `yaml:"-"`
	Value interface{} `yaml:"-"`
}

// WHERE I AM - i can load specs
// now, need to generate a hard cope of everything in /templates
// need to shape the plugin spec to the template, or vice versa.
// latter probably faster that this point unless i need to massage
// some datas.
func main() {
	specLoc := flag.String("spec", "", "The path to the spec file")
	packageRoot := flag.String("package", "", "The go package root for this plugin. Ex: github.com/you/pluginname")
	flag.Parse()
	if *specLoc == "" {
		log.Fatal("Error, must provide path to spec including name")
	}
	if *packageRoot == "" {
		log.Fatal("Error, must provide a package root for the resulting go package. Ex: github.com/you/pluginname")
	}
	data, err := ioutil.ReadFile(*specLoc)
	if err != nil {
		log.Fatalf("error loading spec: %s", err)
	}
	// Chop off the trailing slash, we add it back in as we need it.
	pRoot := *packageRoot
	fmt.Println(pRoot)
	if strings.HasSuffix(*packageRoot, "/") {
		pRoot = pRoot[0 : len(pRoot)-1]
	}
	s := &PluginSpec{
		PackageRoot:  pRoot,
		SpecLocation: *specLoc,
	}
	if err := yaml.Unmarshal(data, s); err != nil {
		log.Fatal(err)
	}
	// Fill in some basic stuff that isn't readily available from the parse, but helps the generation
	postProcessSpec(s)
	// Now, MAKE IT HAPPEN
	if err := generatePlugin(s); err != nil {
		log.Fatal(err)
	}
}

// postProcessSpec does some minor post-processing on the spec object to fill a few things in that make
// template generation easier
func postProcessSpec(s *PluginSpec) {
	// I don't like this dual-dependency shit on typemapper and spec but idgaf right now to bother with it
	t := NewTypeMapper(s)
	s.TypeMapper = t

	// Handle any custom types
	// We'll both populate Types AND update RawTypes so the original source is correct w/r/t the downstream source
	s.Types = make(map[string]TypeData)
	for name, data := range s.RawTypes {
		td := TypeData{}
		td.RawName = name
		td.Name = UpperCamelCase(name)

		for pname, params := range data {
			params.RawName = pname
			params.Name = UpperCamelCase(pname)
			params.Type = t.SpecTypeToGoType(params.Type)
			data[pname] = params
		}
		td.Fields = data
		s.RawTypes[name] = data
		s.Types[name] = td
	}

	// fill in the connection names
	for name, param := range s.Connection {
		param.RawName = name
		param.Name = UpperCamelCase(name)
		param.Type = t.SpecTypeToGoType(param.Type)
		s.Connection[name] = param
	}
	// fill in the trigger names
	for name, action := range s.Actions {
		action.RawName = name // not set in the yaml this way, but set for the benefit of the template
		action.Name = UpperCamelCase(name)
		action.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		for name, param := range action.Input {
			param.RawName = name
			param.Name = UpperCamelCase(name)
			param.Type = t.SpecTypeToGoType(param.Type)
			if strings.Index(param.Type, "types.") > -1 {
				action.NeedsTypes = true
			}
			if strings.Index(param.Type, "time.Time") > -1 {
				action.NeedsTime = true
			}
			action.Input[name] = param // Godbless go for this feature
		}
		for name, param := range action.Output {
			param.RawName = name
			param.Name = UpperCamelCase(name)
			param.Type = t.SpecTypeToGoType(param.Type)
			if strings.Index(param.Type, "types.") > -1 {
				action.NeedsTypes = true
			}
			if strings.Index(param.Type, "time.Time") > -1 {
				action.NeedsTime = true
			}
			action.Output[name] = param // Godbless go for this feature
		}
		s.Actions[name] = action
	}

	for name, trigger := range s.Triggers {
		trigger.RawName = name // not set in the yaml this way, but set for the benefit of the template
		trigger.Name = UpperCamelCase(name)
		trigger.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		for name, param := range trigger.Input {
			param.RawName = name
			param.Name = UpperCamelCase(name)
			param.Type = t.SpecTypeToGoType(param.Type)
			if strings.Index(param.Type, "types.") > -1 {
				trigger.NeedsTypes = true
			}
			if strings.Index(param.Type, "time.Time") > -1 {
				trigger.NeedsTime = true
			}
			trigger.Input[name] = param // Godbless go for this feature
		}
		for name, param := range trigger.Output {
			param.RawName = name
			param.Name = UpperCamelCase(name)
			param.Type = t.SpecTypeToGoType(param.Type)
			if strings.Index(param.Type, "types.") > -1 {
				trigger.NeedsTypes = true
			}
			if strings.Index(param.Type, "time.Time") > -1 {
				trigger.NeedsTime = true
			}
			trigger.Output[name] = param // Godbless go for this feature
		}
		s.Triggers[name] = trigger
	}

	if s.HTTP.Port == 0 {
		s.HTTP.Port = 10001
	}

	if s.HTTP.ReadTimeout == 0 {
		s.HTTP.ReadTimeout = 2
	}

	if s.HTTP.WriteTimeout == 0 {
		s.HTTP.WriteTimeout = 2
	}
}

func generatePlugin(s *PluginSpec) error {
	// Get GOPATH and then add the plugin root
	p := path.Join(os.Getenv("GOPATH"), "src", s.PackageRoot)
	// if it exists, fail
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// if it doesn't, mkdir-p it
		if err = os.MkdirAll(p, 0700); err != nil {
			log.Fatal("Error when creating plugin package path: " + err.Error())
		}
	}
	if err := generateActions(s); err != nil {
		return err
	}
	if err := generateConnections(s); err != nil {
		return err
	}
	if err := generateTriggers(s); err != nil {
		return err
	}
	if err := generateCmd(s); err != nil {
		return err
	}
	if err := generateHTTPServer(s); err != nil {
		return err
	}
	if err := generateHTTPHandlers(s); err != nil {
		return err
	}
	if err := generateTests(s); err != nil {
		return err
	}
	if err := generateTypes(s); err != nil {
		return err
	}
	if err := generateBuildSupport(s); err != nil {
		return err
	}
	if err := copySpec(s); err != nil {
		return err
	}
	// run go FMT before any vendoring
	if err := runGoFMT(s); err != nil {
		return err
	}
	return nil
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err) // return !notexisterror because the function is checking if it DOES exist, not if it DOES NOT
}

func runTemplate(templatePath string, outputPath string, data interface{}, skipIfExists bool) error {
	if skipIfExists && doesFileExist(outputPath) {
		return nil // This isn't error, this is just the function deciding to do nothing under these circumstances
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0700); err != nil {
		return err
	}
	b, err := Asset(templatePath)
	if err != nil {
		return err
	}
	tmp := template.New(templatePath)
	t, err := tmp.Parse(string(b))
	if err != nil {
		return err
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	// now run the template
	if err := t.Execute(f, data); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func generateActions(s *PluginSpec) error {
	// Now, do one for each action using the action_x template
	pathToActionTemplate := "templates/actions/action_x.template"
	pathToRunTemplate := "templates/actions/action_x_run.template"
	for name, action := range s.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/actions/", name+".go")
		if err := runTemplate(pathToActionTemplate, newFilePath, action, false); err != nil {
			return err
		}
		// Make the new action_run.go
		// action_run is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/actions/", name+"_run.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, action, true); err != nil {
			return err
		}
	}
	return nil
}

func generateConnections(s *PluginSpec) error {
	pathToTemplate := "templates/connection/connection.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/connection/", "connection.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}
	// Connect and validate are broken out, so that re-generating will skip them if they exist, making it easier for the dev
	pathToTemplate = "templates/connection/connect.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/connection/connect.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, true); err != nil {
		return err
	}
	pathToTemplate = "templates/connection/cache.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/connection/cache.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}
	pathToTemplate = "templates/connection/connection_key.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/connection/connection_key.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, true); err != nil {
		return err
	}
	return nil
}

func generateTriggers(s *PluginSpec) error {
	// Now, do one for each action using the action_x template
	pathToTriggerTemplate := "templates/triggers/trigger_x.template"
	pathToRunTemplate := "templates/triggers/trigger_x_run.template"

	for name, trigger := range s.Triggers {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/triggers/", name+".go")
		if err := runTemplate(pathToTriggerTemplate, newFilePath, trigger, false); err != nil {
			return err
		}
		// trigger_run is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/triggers/", name+"_run.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, trigger, true); err != nil {
			return err
		}
	}
	return nil
}

func generateCmd(s *PluginSpec) error {
	pathToTemplate := "templates/cmd/main.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/cmd/", "main.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}
	return nil
}

func generateHTTPServer(s *PluginSpec) error {
	pathToTemplate := "templates/server/http/server.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/server/http/", "server.go")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}
	return nil
}

func generateHTTPHandlers(s *PluginSpec) error {
	// Now, do one for each action using the action_x template
	pathToTemplate := "templates/server/http/handler_x.template"
	for name, action := range s.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/server/http/", name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, action, false); err != nil {
			return err
		}
	}
	return nil
}

func generateTests(s *PluginSpec) error {
	return nil
}

func generateTypes(s *PluginSpec) error {
	// First, in case there are no triggers, make a placeholder
	pathToPlaceholderTemplate := "templates/types/types.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/types/types.go")
	if err := runTemplate(pathToPlaceholderTemplate, newFilePath, s, false); err != nil {
		return err
	}
	// Now, do one for each action using the action_x template
	pathToTemplate := "templates/types/type_x.template"
	for name, t := range s.Types {
		// Make the new action.go
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/types/", name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, t, false); err != nil {
			return err
		}
	}
	return nil
}

func generateBuildSupport(s *PluginSpec) error {
	// Docker
	pathToTemplate := "templates/Dockerfile.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "Dockerfile")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}

	// Make
	pathToTemplate = "templates/Makefile.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "Makefile")
	if err := runTemplate(pathToTemplate, newFilePath, s, false); err != nil {
		return err
	}

	// Make the vendor directory for them, add a .gitkeep just incase
	if err := os.MkdirAll(path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/vendor/"), 0700); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "/vendor/.gitkeep"), make([]byte, 0), 0644); err != nil {
		return err
	}
	return nil
}

func copySpec(s *PluginSpec) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(s.SpecLocation)
	if err != nil {
		return err
	}
	// Write data to dst
	return ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, "plugin.spec.yaml"), data, 0644)
}

func runGoFMT(s *PluginSpec) error {
	// TODO add tests?
	// TODO pivot to using goimports, which expects whole files, not packages :/
	vendorablePackages := []string{"actions", "cmd", "connection", "server/http", "triggers", "types"}
	for _, p := range vendorablePackages {
		// Some plugins won't have actions or triggers, so stat the path first
		if doesFileExist(path.Join(os.Getenv("GOPATH"), "/src/", s.PackageRoot, p)) {
			cmd := exec.Command("go", "fmt", s.PackageRoot+"/"+p)
			if b, err := cmd.CombinedOutput(); err != nil {
				fmt.Println(string(b))
				return err
			}
		}
	}
	return nil
}
