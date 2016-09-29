# CODE WE WILL USE FROM THE STANDARD LIBRARY

## [template.Template](https://godoc.org/text/template#Template)
``` Go
template.Template
```

***

## [template.ParseFiles](https://godoc.org/text/template#ParseFiles)
``` Go
func ParseFiles(filenames ...string) (*Template, error)
```

## [template.ParseGlob](https://godoc.org/text/template#ParseGlob)
``` Go
func ParseGlob(pattern string) (*Template, error)
```
***

## [template.Parse](https://godoc.org/text/template#Template.Parse)
``` Go
func (t *Template) Parse(text string) (*Template, error)
```

## [template.ParseFiles](https://godoc.org/text/template#Template.ParseFiles)
``` Go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

## [template.ParseGlob](https://godoc.org/text/template#Template.ParseGlob)
``` Go
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

***

## [template.Execute](https://godoc.org/text/template#Template.Execute)
``` Go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

## [template.ExecuteTemplate](https://godoc.org/text/template#Template.ExecuteTemplate)
``` Go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

***

## [template.Must](https://godoc.org/text/template#Must)
``` Go
func Must(t *Template, err error) *Template
```

## [template.New](https://godoc.org/text/template#New)
``` Go
func New(name string) *Template
```

***

## [The init function](https://golang.org/doc/effective_go.html#init)
``` Go
func init()
```