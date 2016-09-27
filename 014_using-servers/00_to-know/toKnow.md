
- [http.ListenAndServer](https://godoc.org/net/http#ListenAndServe)
-func ListenAndServe(addr string, handler Handler) error

- [http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)
-func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error

- [http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```