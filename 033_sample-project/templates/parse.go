package templates

import (
	"html/template"
	"io"
)

// TPL holds all parsed templates
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func (t *template.Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t.ExecuteTemplate
}
