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

	yaml "gopkg.in/yaml.v2"
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
		Actions:      make(map[string]PluginHandlerData),
		Triggers:     make(map[string]PluginHandlerData),
		Connection:   make(map[string]ParamData),
	}
	if err = yaml.Unmarshal(data, s); err != nil {
		return nil, err
	}
	// Fill in some basic stuff that isn't readily available from the parse, but helps the generation
	err = PostProcessSpec(s)
	return &Generator{
		spec:          s,
		forceOverwise: force,
	}, err
}

// PluginSpec is a spec for plugins
type PluginSpec struct {
	PluginSpecVersion             string                         `yaml:"plugin_spec_version"`
	Name                          string                         `yaml:"name"`
	Title                         string                         `yaml:"title"`
	Description                   string                         `yaml:"description"`
	Version                       string                         `yaml:"version"`
	Vendor                        string                         `yaml:"vendor"`
	Tags                          []string                       `yaml:"tags"`
	Icon                          string                         `yaml:"icon"`
	Help                          string                         `yaml:"help"`
	Connection                    ParamDataCollection            `yaml:"connection"`
	ConnectionRequiresCustomTypes bool                           `yaml:"-"`
	ConnectionSchema              JSONSchema                     `yaml:"-"`
	RawTypes                      map[string]ParamDataCollection `yaml:"types"`
	Types                         map[string]TypeData            `yaml:"-"`
	Triggers                      PluginHandlerCollection        `yaml:"triggers"`
	Actions                       PluginHandlerCollection        `yaml:"actions"`
	// Things that are not part of the spec, but still important
	PackageRoot string `yaml:"package_root"`
	HTTP        HTTP   `yaml:"http"`

	// Things that are for background help
	TypeMapper        *TypeMapper `yaml:"-"`
	SpecLocation      string      `yaml:"-"`
	ConnectionDataKey string      `yaml:"-"`
}

// TypeData defines the custom types. Much of the data is pulled via a parent-key, so we don't parse much from yaml at all.
// Instead, we post-process populate it for the benefit of the template
type TypeData struct {
	RawName      string              `yaml:"-"`
	Name         string              `yaml:"-"`
	Fields       ParamDataCollection `yaml:"-"`
	SortedFields []ParamData         `yaml:"-"`
	Schema       JSONSchema
}

// HTTP Defines the settings for the plugins http server
type HTTP struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
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
	return g.vendorPluginDeps()
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
	fmt.Printf("Generating %s/actions\n", g.spec.PackageRoot)
	// Now, do one for each action using the action_x template
	pathToActionTemplate := "templates/actions/action_x.template"
	pathToRunTemplate := "templates/actions/action_x_custom.template"
	for name, action := range g.spec.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/actions/", name+".go")
		if err := runTemplate(pathToActionTemplate, newFilePath, action, false); err != nil {
			return err
		}
		// Make the new action_custom.go
		// action_custom is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/actions/", name+"_custom.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, action, true); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateConnections() error {
	fmt.Printf("Generating %s/connection\n", g.spec.PackageRoot)
	pathToTemplate := "templates/connection/connection.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/", "connection.go")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
		return err
	}
	// connection_custom and validate are broken out, so that re-generating will skip them if they exist, making it easier for the dev
	pathToTemplate = "templates/connection/connection_custom.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/connection_custom.go")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, true); err != nil {
		return err
	}
	pathToTemplate = "templates/connection/connection_tester.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/connection_tester.go")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, true); err != nil {
		return err
	}
	pathToTemplate = "templates/connection/cache.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/connection/cache.go")
	return runTemplate(pathToTemplate, newFilePath, nil, false)
}

func (g *Generator) generateTriggers() error {
	fmt.Printf("Generating %s/triggers\n", g.spec.PackageRoot)
	// Now, do one for each action using the action_x template
	pathToTriggerTemplate := "templates/triggers/trigger_x.template"
	pathToRunTemplate := "templates/triggers/trigger_x_custom.template"

	for name, trigger := range g.spec.Triggers {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/triggers/", name+".go")
		if err := runTemplate(pathToTriggerTemplate, newFilePath, trigger, false); err != nil {
			return err
		}
		// trigger_run is broken out, so that re-generating will skip them if they exist, making it easier for the dev
		newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/triggers/", name+"_custom.go")
		if err := runTemplate(pathToRunTemplate, newFilePath, trigger, true); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateCmd() error {
	fmt.Printf("Generating %s/cmd\n", g.spec.PackageRoot)
	pathToTemplate := "templates/cmd/main.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/cmd/", "main.go")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) generateHTTPServer() error {
	fmt.Printf("Generating %s/server/http (daemon)\n", g.spec.PackageRoot)
	pathToTemplate := "templates/server/http/server.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/server/http/", "server.go")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) generateHTTPHandlers() error {
	fmt.Printf("Generating %s/server/http (handlers)\n", g.spec.PackageRoot)
	// Now, do one for each action using the action_x template
	pathToTemplate := "templates/server/http/handler_x.template"
	for name, action := range g.spec.Actions {
		// Make the new action.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/server/http/", name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, action, false); err != nil {
			return err
		}
	}
	pathToTemplate = "templates/server/http/handler_trigger_x.template"
	for name, action := range g.spec.Triggers {
		// Make the new trigger.go
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/server/http/", "trigger_"+name+".go")
		if err := runTemplate(pathToTemplate, newFilePath, action, false); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateTests() error {
	fmt.Printf("Generating %s/tests\n", g.spec.PackageRoot)
	return nil
}

func (g *Generator) generateTypes() error {
	fmt.Printf("Generating %s/types\n", g.spec.PackageRoot)
	tList := []string{
		"sdk_file",
		"credential_asymmetric_key",
		"credential_token",
		"credential_username_password",
		"credential_secret_key",
	}
	// Do the stock ones first
	for _, t := range tList {
		pathToTemplate := "templates/types/" + t + ".template"
		newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "/types/"+t+".go")
		if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
			return err
		}
	}
	// Now, do one for each type using the type_x template
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
	fmt.Printf("Generating %s Build Artifacts (make, docker) \n", g.spec.PackageRoot)
	// Docker
	pathToTemplate := "templates/Dockerfile.template"
	newFilePath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "Dockerfile")
	if err := runTemplate(pathToTemplate, newFilePath, g.spec, false); err != nil {
		return err
	}

	// Make
	pathToTemplate = "templates/Makefile.template"
	newFilePath = path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "Makefile")
	return runTemplate(pathToTemplate, newFilePath, g.spec, false)
}

func (g *Generator) copySpec() error {
	fmt.Printf("Generating %s/plugin.spec.yaml\n", g.spec.PackageRoot)
	// Read all content of src to data
	data, err := ioutil.ReadFile(g.spec.SpecLocation)
	if err != nil {
		return err
	}
	// Write data to dst
	return ioutil.WriteFile(path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot, "plugin.spec.yaml"), data, 0644)
}

func (g *Generator) runGoImports() error {
	fmt.Println("Formatting and updating import statements")
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
			return fmt.Errorf("error while running go imports on %s: %s", p, string(b))
		}
	}
	return nil
}

func (g *Generator) vendorPluginDeps() error {
	fmt.Println("Vendoring dependencies... ")
	fmt.Println(" ...This may take a few minutes depending on your network connection")
	rootPath := path.Join(os.Getenv("GOPATH"), "/src/", g.spec.PackageRoot)
	if doesFileExist(path.Join(rootPath, "vendor")) {
		// run ensure
		cmd := exec.Command("dep", "ensure", "-v") // nolint: gas
		cmd.Dir = rootPath
		if b, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("error while running dep ensure on %s: %s - %s", rootPath, string(b), err.Error())
		}
	} else {
		// first time, run init
		cmd := exec.Command("dep", "init", "-v") // nolint: gas
		cmd.Dir = rootPath
		if b, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("error while running dep init on %s: %s - %s", rootPath, string(b), err.Error())
		}
	}
	return nil
}
