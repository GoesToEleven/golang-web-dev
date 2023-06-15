# Web Programming Synonymous Terms

- router
- request router
- multiplexer
- mux
- servemux
- server
- http router
- http request router
- http multiplexer
- http mux
- http servemux
- http server

---

In electronics, a [multiplexer (or mux)](https://en.wikipedia.org/wiki/Multiplexer) is a device that selects one of several input signals and forwards the selected input into a single line.

The term multiplexer has been adopted by web programming to refer to the process of routing requests.

A web server has requests coming in at different routers and via different HTTP methods. For instance, we might have these requests:

REQUEST #1

- Path: /cat
- Method: GET

REQUEST #2

- Path: /apply
- Method: Get

Request #3

- Path: /apply
- Method: Post

Based upon the requests coming in, the server needs to determine how to respond to that request - for each request that comes in, different code will be run.

I've been using the word "server" but I could have also been using the word "multiplexer" or "mux". The server, or multiplexer, or mux, determines what code needs to be run in response to each incoming request

---

ServeMux is an HTTP request multiplexer.

A ServeMux matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash).

Longer patterns take precedence over shorter ones, so that if there are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will be called for paths beginning "/images/thumbnails/" and the former will receive requests for any other paths in the "/images/" subtree. Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".

If a subtree has been registered and a request is received naming the subtree root without its trailing slash, ServeMux redirects that request to the subtree root (adding the trailing slash). This behavior can be overridden with a separate registration for the path without the trailing slash. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately.

Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/".

ServeMux also takes care of sanitizing the URL request path, redirecting any request containing . or .. elements or repeated slashes to an equivalent, cleaner URL.

---

# ServeMux

[http.ServeMux](https://godoc.org/net/http#ServeMux)

```Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

Any value of type `*http.ServeMux` implements the `http.Handler` interface.

Remember, the `http.Handler` interface requires that a type have the `ServeHTTP` method.

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

What this tells us is that we can pass a value of type `*http.ServeMux` into `http.ListenAndServe`

```
func ListenAndServe(addr string, handler Handler) error
```

You can also see from the documentation of `*http.ServeMux` ...

```Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

... that we have a method `Handle` attached the the value of type `*http.ServeMux`. You can see that `Handle` takes a pattern and a `http.Handler`.

We can use `Handle` like this:

```
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/cat", c)
```

The overall game plan:

We will create a NewServeMux, then attach the method `Handle` to it to set routes, then pass our `*http.ServeMux` to `http.ListenAndServe`.

---

# DefaultServeMux

ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

```
http.ListenAndServe(":8080", nil)
```

---

Read https://pkg.go.dev/net/http

## func Handle

- func Handle은 두 번째 매개변수에 Handler를 바로 사용한다.

```Go


func Handle(pattern string, handler Handler)
Handle registers the handler for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.



import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}



```

## func HandleFunc

- HandleFunc은 하위에 ResponseWriter와 Request를 가지고 있는 함수를 바로 사용한다.

```Go

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
HandleFunc registers the handler function for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.


package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


```

## func ListenAndServe

```Go

func ListenAndServe(addr string, handler Handler) error
ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

The handler is typically nil, in which case the DefaultServeMux is used.

ListenAndServe always returns a non-nil error.


package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}




```

## type HandlerFunc

```Go
type HandlerFunc func(ResponseWriter, *Request)
```

The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

## func (HandlerFunc) ServeHTTP

```Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

```

ServeHTTP calls f(w, r).
