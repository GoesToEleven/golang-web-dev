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

***

# Request & response

Request and response messages are similar. Both messages consist of:

- a request/response line
- zero or more header lines
- a blank line (ie, CRLF by itself)
- an optional message body

***

## HTTP request

Request
- request line
- headers
- optional message body

Request-Line
- Method SP Request-URI SP HTTP-Version CRLF

example request line:
- GET /path/to/file/index.html HTTP/1.1

***

## HTTP response

Reponse
- status line
- headers
- optional message body

Status-Line
- HTTP-Version SP Status-Code SP Reason-Phrase CRLF

example status line:
- HTTP/1.1 200 OK

***

## Headers
[List of header fields](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)

***

## Inspect
- you can use google chrome dev tools / network
- you can use CURL at the command line

```
curl -v golang.org
```
