# HTTP/1.1 message

- a start-line 
- zero or more header fields 
- an empty line indicating the end of the header section 
- an optional message body.

***

## Request line (request)

GET / HTTP/1.1

[method SP request-target SP HTTP-version CRLF](https://tools.ietf.org/html/rfc7230#section-3.1.1)

## Status line (response)

HTTP/1.1 302 Found

[HTTP-version SP status-code SP reason-phrase CRLF](https://tools.ietf.org/html/rfc7230#section-3.1.2)

***

# Writing a response

``` Go
body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
io.WriteString(conn, "HTTP/1.1 200 OK\r\n") // status line
fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) // header
fmt.Fprint(conn, "Content-Type: text/plain\r\n") // header
io.WriteString(conn, "\r\n") // blank line; CRLF; carriage-return line-feed
io.WriteString(conn, body) // body, aka, payload
```

***

# Parsing String

[strings.Fields](https://godoc.org/strings#Fields)
``` Go
func Fields(s string) []string
```