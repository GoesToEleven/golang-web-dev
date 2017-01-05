# Nested templates

[nested templates documentation](https://godoc.org/text/template#hdr-Nested_template_definitions)

## define: 
``` Go
{{define "TemplateName"}}
insert content here
{{end}}
```
## use: 
``` Go
{{template "TemplateName"}}
```