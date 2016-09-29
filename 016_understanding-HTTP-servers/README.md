# Request line (request)

GET / HTTP/1.1

method SP request-target SP HTTP-version CRLF

https://tools.ietf.org/html/rfc7230#section-3.1.1

# Status line (response)

HTTP/1.1 302 Found

HTTP-version SP status-code SP reason-phrase CRLF

https://tools.ietf.org/html/rfc7230#section-3.1.2

***

## Writing a response

``` Go
body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
fmt.Fprint(conn, "Content-Type: text/plain\r\n")
io.WriteString(conn, "\r\n")
io.WriteString(conn, body)
```

***

All HTTP/1.1 messages consist of a start-line 
followed by a sequence of octets in a format similar to 
the Internet Message Format: zero or more header fields 
(collectively referred to as the "headers" or the "header section"), 
an empty line indicating the end of the header section, 
and an optional message body.

***

# Parsing String
- [strings.Fields](https://godoc.org/strings#Fields)
``` Go
func Fields(s string) []string
```