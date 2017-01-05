# Template variables

## [template variables](https://godoc.org/text/template#hdr-Variables)

### ASSIGN
``` Go
{{$wisdom := .}}
```

### USE
``` Go
{{$wisdom}}
```

A pipeline inside an action may initialize a variable to capture the result. The initialization has syntax
 
 $variable := pipeline
 
 where $variable is the name of the variable. An action that declares a variable produces no output.
 
 If a "range" action initializes a variable, the variable is set to the successive elements of the iteration. Also, a "range" may declare two variables, separated by a comma:
 
  range $index, $element := pipeline
  
 in which case $index and $element are set to the successive values of the array/slice index or map key and element, respectively. Note that if there is only one variable, it is assigned the element; this is opposite to the convention in Go range clauses.
 
 A variable's scope extends to the "end" action of the control structure ("if", "with", or "range") in which it is declared, or to the end of the template if there is no such control structure. A template invocation does not inherit variables from the point of its invocation.
 
 When execution begins, $ is set to the data argument passed to Execute, that is, to the starting value of dot.