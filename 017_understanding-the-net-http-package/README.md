# Handler

*This is one of the most important things to know*

[http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

***

# Server

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)
``` Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

*Notice that both of the above functions take a handler*

***

# Request

See [http.Request](https://godoc.org/net/http#Request) in the documentation.
 
Here it is with *most of the comments and some of the fields* stripped out:

```go 
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}
```

Also see the index showing type [Request]() from the http package.

Some interesting things you can do with a request:

## Retrieve URL & Form data

```http.Request``` is a struct. It has the fields ```Form``` & ```PostForm```. If we read the documentation on these, we see:

```
    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values

```

If we look at **ParseForm**, 

```go func (r *Request) ParseForm() error ```

we see that this is a method attached to a *http.Request.

***

If we look at **FormValue***

``` go func (r *Request) FormValue(key string) string```

we see that this is a method attached to a *http.Request. FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.


***

## See the HTTP Method

The ```http.Request``` type is a struct which has a ```Method``` field.

***

## See URL values

The ```http.Request``` type is a struct which has a ```URL``` field. Notice that the type is a ```*url.URL```

Take a look at type ```url.URL```

``` go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```

***

## Work with the HTTP header

The ```http.Request``` type is a struct which has a ```Header``` field. 

From the documentation about the ```http.Header``` struct, we see that:

```
type Header map[string][]string
```

***

# Response

[http.ResponseWriter](https://godoc.org/net/http#ResponseWriter)
``` Go
type ResponseWriter interface {
    // Header returns the header map that will be sent by
    // WriteHeader. Changing the header after a call to
    // WriteHeader (or Write) has no effect 
    Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
    //
    // If WriteHeader has not yet been called, Write calls
    // WriteHeader(http.StatusOK) before writing the data. If the Header
    // does not contain a Content-Type line, Write adds a Content-Type set
    // to the result of passing the initial 512 bytes of written data to
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with status code.
    // If WriteHeader is not called explicitly, the first call to Write
    // will trigger an implicit WriteHeader(http.StatusOK).
    // Thus explicit calls to WriteHeader are mainly used to
    // send error codes.
    WriteHeader(int)
}
```

***

## Setting a response header

An ```http.ResponseWriter``` has a method ```Header()``` which returns a ```http.Header```.

Look at the documentation for ```http.Header```

``` Go
type Header map[string][]string

```

Look at the methods which are attached to type ```http.Header```

``` go
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

We can set headers for a response like this:

``` Go
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```