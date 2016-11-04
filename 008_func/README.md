# Using functions in templates

## [template function documentation](https://godoc.org/text/template#hdr-Functions)

***

## [template.FuncMap](type FuncMap map[string]interface{})

FuncMap is the type of the map defining the mapping from names to functions. Each function must have either a single return value, or two return values of which the second has type error. In that case, if the second (error) return value evaluates to non-nil during execution, execution terminates and Execute returns that error.

## [template.Funcs](https://godoc.org/text/template#Template.Funcs)
``` Go
func (t *Template) Funcs(funcMap FuncMap) *Template
```

***

During execution functions are found in two function maps: 
- first in the template, 
- then in the global function map. 

By default, no functions are defined in the template but the Funcs method can be used to add them.

Predefined global functions are defined in text/template.