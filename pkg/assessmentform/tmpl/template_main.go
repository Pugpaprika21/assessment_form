package tmpl

import (
	"errors"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
)

type (
	templateRegistry struct {
		templates         map[string]*template.Template
		templatesRegister string
	}
	templatesConfig struct {
		Templates map[string][]string `yaml:"templates"`
	}
)

func NewTemplateRegistry() *templateRegistry {
	return &templateRegistry{
		templates: make(map[string]*template.Template),
		/* templates_register.yaml */
		templatesRegister: "pkg/service-name/tmpl/templates_register.yaml",
	}
}

func (t *templateRegistry) Load() *templateRegistry {
	data, err := os.ReadFile(t.templatesRegister)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config templatesConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}

	for name, files := range config.Templates {
		t.templates[name] = template.Must(template.ParseFiles(files...))
	}

	return t
}

func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return errors.New("template not found -> " + name)
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
