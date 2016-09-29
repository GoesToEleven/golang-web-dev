# ServeMux

[http.ServeMux](https://godoc.org/net/http#ServeMux)

``` Go
type ServeMux struct {
    // contains filtered or unexported fields
}
```

ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL. Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash). Longer patterns take precedence over shorter ones, so that if there are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will be called for paths beginning "/images/thumbnails/" and the former will receive requests for any other paths in the "/images/" subtree. Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/". 

If a subtree has been registered and a request is received naming the subtree root without its trailing slash, ServeMux redirects that request to the subtree root (adding the trailing slash). This behavior can be overridden with a separate registration for the path without the trailing slash. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately. 

Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/". 

ServeMux also takes care of sanitizing the URL request path, redirecting any request containing . or .. elements or repeated slashes to an equivalent, cleaner URL.

## [http.NewServeMux](https://godoc.org/net/http#NewServeMux)
``` Go
func NewServeMux() *ServeMux
```

## [http.Handle](https://godoc.org/net/http#ServeMux.Handle)
``` Go
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

## [http.ServeHTTP](https://godoc.org/net/http#ServeMux.ServeHTTP)
``` Go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

As *ServeMux has the ServeHTTP method, it implicitly implements the Handler interface. A *ServeMux is also of type Handler. This means that we can pass a *ServeMux to ListenAndServe.

***

# Review

## Handler

*This is one of the most important things to know!*

[http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

***

## Server

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)
``` Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

***