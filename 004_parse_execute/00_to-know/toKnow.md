``` Go
template.Template
```
***********************************
``` Go
func ParseFiles(filenames ...string) (*Template, error)

func ParseGlob(pattern string) (*Template, error)
```
***********************************
``` Go
func (t *Template) Parse(text string) (*Template, error)

func (t *Template) ParseFiles(filenames ...string) (*Template, error)

func (t *Template) ParseGlob(pattern string) (*Template, error)
```
***********************************
``` Go
func (t *Template) Execute(wr io.Writer, data interface{}) error

func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```
***********************************
``` Go
func Must(t *Template, err error) *Template

func New(name string) *Template
```
***********************************
``` Go
func init()
```