look at the connection: REQUEST & RESPONSE

- you can use google chrome dev tools / network
- you can use CURL at the command line
	curl -v golang.org

Specification

REQUEST
GET / HTTP/1.1
method SP request-target SP HTTP-version CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.1

RESPONSE (STATUS)
HTTP/1.1 302 Found
HTTP-version SP status-code SP reason-phrase CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.2