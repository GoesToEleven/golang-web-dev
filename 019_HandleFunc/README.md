# HandleFunc & HandlerFunc

## HandleFunc

[http.HandleFunc](https://godoc.org/net/http#HandleFunc)
``` Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

## HandlerFunc

``` Go
type HandlerFunc func(ResponseWriter, *Request)
```

``` Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

***

## QUESTION
Could you get http.Handle to take a func with this signature: func(ResponseWriter, *Request)?