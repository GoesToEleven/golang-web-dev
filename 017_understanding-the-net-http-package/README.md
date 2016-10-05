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

Also see the index showing type [Request]() from the http package.

Some interesting things you can do with a request:

## Retrieve URL & Form data - all of it

```http.Request``` is a struct. It has the fields ```Form``` & ```PostForm```. If we read the documentation on these, we see:

```
    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    //
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values

```

If we look at **ParseForm**, 


### func (*Request) ParseForm
```go func (r *Request) ParseForm() error ```

we see that this is a method attached to a *http.Request.

##  Retrieve URL & Form data - by identifier (key)


### func (*Request) FormValue
``` gofunc (r *Request) FormValue(key string) string```
FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.


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

## Request URL

req.URL.RequestURI()

req.URL.Path

req.URL

Example:
``` Go
switch req.URL.Path {
	case "/cat":
		io.WriteString(res, `some text`)
	case "/dog":
		io.WriteString(res, `some other text`)
	}
```

***

## Setting a response header

Example:
``` Go
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```
