package modgen

import (
	"bytes"
	"html/template"
	"os"
	"path"
)

// Generator contains the context required to generate a Go module import URL static site.
type Generator struct {
	config     Config
	indexTmpl  *template.Template
	moduleTmpl *template.Template
}

func NewGenerator(config Config) *Generator {
	return &Generator{
		config:     config,
		indexTmpl:  indexTmpl,
		moduleTmpl: moduleTmpl,
	}
}

func (g *Generator) WithIndexTemplate(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tmpl, err := template.New("index-custom").Parse(string(raw))
	if err != nil {
		return err
	}
	g.indexTmpl = tmpl

	return nil
}

func (g *Generator) WithModuleTemplate(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tmpl, err := template.New("module-custom").Parse(string(raw))
	if err != nil {
		return err
	}
	g.moduleTmpl = tmpl

	return nil
}

func (g *Generator) Generate(target string) error {
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	if err := g.generateIndex(path.Join(target, "index.html")); err != nil {
		return err
	}

	for _, module := range g.config.Modules {
		modulePath := path.Join(target, module.Path+".html")
		if err := g.generateModule(modulePath, module); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateIndex(target string) error {
	b := bytes.Buffer{}
	if err := g.indexTmpl.Execute(&b, g.config); err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), 0644)
}

func (g *Generator) generateModule(target string, module ModuleConfig) error {
	b := bytes.Buffer{}
	if err := g.moduleTmpl.Execute(&b, module.withHost(g.config.Host)); err != nil {
		return err
	}

	return os.WriteFile(target, b.Bytes(), 0644)
}
