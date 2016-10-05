# Server

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)
``` Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

***

# Handler

## This is one of the most important things to know!

[http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

***

# Request

[http.Request](https://godoc.org/net/http#Request)
``` Go 
type Request struct {
    // Method specifies the HTTP method (GET, POST, PUT, etc.).
    // For client requests an empty string means GET.
    Method string

    // URL specifies either the URI being requested (for server
    // requests) or the URL to access (for client requests).
    URL *url.URL

    // Header contains the request header fields either received
    // by the server or to be sent by the client.
    Header Header

    // Body is the request's body.
    Body io.ReadCloser

	// other methods not mentioned
}
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
