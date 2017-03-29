package sdk

import (
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

// Generator is a wrapper around building a plugin based on a spec
type Generator struct {
	spec          *PluginSpec
	forceOverwise bool
}

// NewGenerator takes a path to a spec, and the root of the go package, and spins up a new plugin with all the fixins'
// specLoc is the location of the soec file on disk, packageroot name of the package, relative to a proper GOPATH
// and force will force the plugin to overwrite any existing files, including ones that don't normally get overwritten in a
// generation, such as the user-editable files
func NewGenerator(specPath, packageRoot string, force bool) (*Generator, error) {
	data, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("error loading spec: %s", err.Error())
	}
	if strings.HasSuffix(packageRoot, "/") {
		packageRoot = packageRoot[0 : len(packageRoot)-1]
	}
	s := &PluginSpec{
		PackageRoot:  packageRoot,
		SpecLocation: specPath,
	}
	if err := yaml.Unmarshal(data, s); err != nil {
		return nil, err
	}
	// Fill in some basic stuff that isn't readily available from the parse, but helps the generation
	postProcessSpec(s)
	return &Generator{
		spec:          s,
		forceOverwise: force,
	}, nil
}

// PluginSpec is a spec for plugins
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
	TypeMapper        *TypeMapper `yaml:"-"`
	SpecLocation      string      `yaml:"-"`
	ConnectionDataKey string      `yaml:"-"`
}

// ParamData are the details that make up each input/output/trigger param
// Not all data is always filled in, it's context sensitive, which is unfortunate but
// I'm willing to accept it for this since everything else is greatly simplified
type ParamData struct {
	RawName     string        `yaml:"name"`
	Name        string        `yaml:"-"` // This is the joined and camelled name for the param
	Type        string        `yaml:"type"`
	Required    bool          `yaml:"required"`
	Description string        `yaml:"description"`
	Enum        []interface{} `yaml:"enum"`
	Default     interface{}   `yaml:"default"`
	Embed       bool          `yaml:"embed"`
	Nullable    bool          `yaml:"nullable"`

	// Things that are used for background help
	EnumLiteral []EnumData `yaml:"-"`
}

// EnumData is used to parse and write out enums
type EnumData struct {
	Name         string `yaml:"-"`
	LiteralValue string `yaml:"-"`
}

// TypesInternalType is a special hack for the types package. Because in all other places we need to reference X via types.X
// we prefix it ahead of time. But types that use or refer to other types in the types package don't need it.
// This method, used only from the code generators for types, is to make sure types don't use the pacakge name internally in the package
func (p ParamData) TypesInternalType() string {
	if i := strings.Index(p.Type, "types."); i > -1 {
		return strings.Replace(p.Type, "types.", "", -1)
	}
	return p.Type
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
}

// TypeData defines the custom types. Much of the data is pulled via a parent-key, so we don't parse much from yaml at all.
// Instead, we post-process populate it for the benefit of the template
type TypeData struct {
	RawName      string               `yaml:"-"`
	Name         string               `yaml:"-"`
	Fields       map[string]ParamData `yaml:"-"`
	SortedFields []ParamData          `yaml:"-"`
}

// HTTP Defines the settings for the plugins http server
type HTTP struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
}

// postProcessSpec does some minor post-processing on the spec object to fill a few things in that make
// template generation easier
func postProcessSpec(s *PluginSpec) error {
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

		if err := updateParams(data, t); err != nil {
			return err
		}

		td.Fields = data
		// Sort them  - currently this is by their embedded status, as embeds must appear uptop in go structs
		td.SortedFields = sortParamData(td.Fields)
		s.RawTypes[name] = data
		s.Types[name] = td
	}

	// fill in the connection names
	// Do this one out long form since we need the special case of building the data key
	for name, param := range s.Connection {
		param.RawName = name
		param.Name = UpperCamelCase(name)
		param.Type = t.SpecTypeToGoType(param.Type)
		s.Connection[name] = param
		if param.Type == "string" {
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
	// fill in the trigger names
	for name, action := range s.Actions {
		action.RawName = name // not set in the yaml this way, but set for the benefit of the template
		action.Name = UpperCamelCase(name)
		action.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		if err := updateParams(action.Input, t); err != nil {
			return err
		}
		if err := updateParams(action.Output, t); err != nil {
			return err
		}
		s.Actions[name] = action
	}

	for name, trigger := range s.Triggers {
		trigger.RawName = name // not set in the yaml this way, but set for the benefit of the template
		trigger.Name = UpperCamelCase(name)
		trigger.PackageRoot = s.PackageRoot
		// We need to do the same thing for the params too
		if err := updateParams(trigger.Input, t); err != nil {
			return err
		}
		if err := updateParams(trigger.Output, t); err != nil {
			return err
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
	return nil
}

// GeneratePlugin emits the new plugin based on the spec to disk
func (g *Generator) GeneratePlugin() error {
	// Get GOPATH and then add the plugin root
	p := path.Join(os.Getenv("GOPATH"), "src", g.spec.PackageRoot)
	// if it exists, fail
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// if it doesn't, mkdir-p it
		if err = os.MkdirAll(p, 0700); err != nil {
			log.Fatal("Error when creating plugin package path: " + err.Error())
		}
	}
	if err := g.generateActions(); err != nil {
		return err
	}
	if err := g.generateConnections(); err != nil {
		return err
	}
	if err := g.generateTriggers(); err != nil {
		return err
	}
	if err := g.generateCmd(); err != nil {
		return err
	}
	if err := g.generateHTTPServer(); err != nil {
		return err
	}
	if err := g.generateHTTPHandlers(); err != nil {
		return err
	}
	if err := g.generateTests(); err != nil {
		return err
	}
	if err := g.generateTypes(); err != nil {
		return err
	}
	if err := g.generateBuildSupport(); err != nil {
		return err
	}
	if err := g.copySpec(); err != nil {
		return err
	}
	// run goimports before any vendoring
	if err := g.runGoImports(); err != nil {
		return err
	}
	if err := g.vendorPluginDeps(); err != nil {
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

func (g *Generator) generateActions() error {
	// Now, do one for each action using the action_x template
	pathToActionTemplate := "templates/actions/action_x.template"
	pathToRunTemplate := "templates/actions/action_x_run.template"
	for name, action := range g.spec.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/actions/", name+".go")
		if err := runTemplate(pathToActionTemplate, newFilePath, action, false); err != nil {
			return err
		}
		// Make the new action_run.go
		// action_run is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/actions/", name+"_run.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, action, true); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateConnections() error {
	pathToTemplate := "templates/connection/connection.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/", "connection.go")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
		return err
	}
	// Connect and validate are broken out, so that re-generating will skip them if they exist, making it easier for the dev
	pathToTemplate = "templates/connection/connect.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/connect.go")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, true); err != nil {
		return err
	}
	pathToTemplate = "templates/connection/cache.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/cache.go")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) generateTriggers() error {
	// Now, do one for each action using the action_x template
	pathToTriggerTemplate := "templates/triggers/trigger_x.template"
	pathToRunTemplate := "templates/triggers/trigger_x_run.template"

	for name, trigger := range g.spec.Triggers {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/triggers/", name+".go")
		if err := runTemplate(pathToTriggerTemplate, newFilePath, trigger, false); err != nil {
			return err
		}
		// trigger_run is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/triggers/", name+"_run.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, trigger, true); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateCmd() error {
	pathToTemplate := "templates/cmd/main.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/cmd/", "main.go")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) generateHTTPServer() error {
	pathToTemplate := "templates/server/http/server.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/server/http/", "server.go")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) generateHTTPHandlers() error {
	// Now, do one for each action using the action_x template
	pathToTemplate := "templates/server/http/handler_x.template"
	for name, action := range g.spec.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/server/http/", name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, action, false); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateTests() error {
	return nil
}

func (g *Generator) generateTypes() error {
	// Now, do one for each action using the type_x template
	pathToTemplate := "templates/types/type_x.template"
	for name, t := range g.spec.Types {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/types/", name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, t, false); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateBuildSupport() error {
	// Docker
	pathToTemplate := "templates/Dockerfile.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "Dockerfile")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
		return err
	}

	// Make
	pathToTemplate = "templates/Makefile.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "Makefile")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
		return err
	}

	// Make the vendor directory for them, add a .gitkeep just incase
	if err := os.MkdirAll(path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/vendor/"), 0700); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/vendor/.gitkeep"), make([]byte, 0), 0644); err != nil {
		return err
	}
	return nil
}

func (g *Generator) copySpec() error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(g.spec.SpecLocation)
	if err != nil {
		return err
	}
	// Write data to dst
	return ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "plugin.spec.yaml"), data, 0644)
}

func (g *Generator) runGoImports() error {
	// TODO add tests?
	// TODO pivot to using goimports, which expects whole files, not packages :/
	searchDir := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot)

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !strings.Contains(path, "/vendor/") && strings.HasSuffix(path, ".go") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	for _, p := range fileList {
		cmd := exec.Command("goimports", "-w", "-srcdir", g.spec.PackageRoot, p)
		if b, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("Error while running go imports on %s: %s", p, string(b))
		}
		if err := g.fixGoImportsNotKnowingHowToLookInLocalVendorFirst(p); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) vendorPluginDeps() error {
	rootPath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot)
	cmd := exec.Command("gvt", "fetch", "github.com/komand/plugin-sdk-go2")
	cmd.Dir = rootPath
	fmt.Println(rootPath)
	if doesFileExist(path.Join(rootPath, "vendor", "github.com", "komand", "plugin-sdk-go2")) {
		cmd = exec.Command("gvt", "update", "github.com/komand/plugin-sdk-go2")
	}
	if b, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("Error while running gvt on %s: %s - %s", rootPath, string(b), err.Error())
	}
	return nil
}

func (g *Generator) fixGoImportsNotKnowingHowToLookInLocalVendorFirst(path string) error {
	old := "github.com/komand/komand/plugins/v1/types"
	new := g.spec.PackageRoot + "/types"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return ioutil.WriteFile(path, []byte(strings.Replace(string(b), old, new, -1)), 0)
}
