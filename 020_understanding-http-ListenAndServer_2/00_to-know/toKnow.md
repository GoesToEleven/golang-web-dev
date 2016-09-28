
- [http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

********************
# HANDLER

## This is one of the most important things to know!

- [http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

********************
# SERVEMUX

- [http.NewServeMux](https://godoc.org/net/http#NewServeMux)
``` Go
func NewServeMux() *ServeMux
```

- [http.ServeHTTP](https://godoc.org/net/http#ServeMux.ServeHTTP)
``` Go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

********************
# HANDLERFUNC

## This is also one of the most important things to know!

- [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
``` Go
type HandlerFunc func(ResponseWriter, *Request)
```

- [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
``` Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
