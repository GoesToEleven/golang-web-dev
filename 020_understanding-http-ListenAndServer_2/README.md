# HandleFunc & HandlerFunc

## HandleFunc

[http.HandleFunc](https://godoc.org/net/http#HandleFunc)
``` Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

## HandlerFunc

## [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
``` Go
type HandlerFunc func(ResponseWriter, *Request)
```

## [http.ServeHTTP](https://godoc.org/net/http#HandlerFunc.ServeHTTP)
``` Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

Note the underlying type of HandlerFunc vs. Handler

***

# Review

## Handler

[http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Handle

- [http.Handle](https://godoc.org/net/http#Handle)
``` Go
func Handle(pattern string, handler Handler)
```

***

## ListenAndServe

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

***

# ServeMux

[http.NewServeMux](https://godoc.org/net/http#NewServeMux)
``` Go
func NewServeMux() *ServeMux
```

[http.ServeHTTP](https://godoc.org/net/http#ServeMux.ServeHTTP)
``` Go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

***

## QUESTION

Could you get http.Handle to take a func with this signature: func(ResponseWriter, *Request)?