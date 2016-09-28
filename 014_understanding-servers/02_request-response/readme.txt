REQUEST & RESPONSE

Request and response messages are similar. Both messages consist of:

- a request/response line
- zero or more header lines
- a blank line (ie, CRLF by itself)
- an optional message body

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

HTTP REQUEST

request line
headers
optional message body

Request-Line
Method SP Request-URI SP HTTP-Version CRLF

example:
GET /path/to/file/index.html HTTP/1.0

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

HTTP RESPONSE

Response-line (Status-Line)
HTTP-Version SP Status-Code SP Reason-Phrase CRLF

examples:
HTTP/1.0 200 OK
HTTP/1.0 404 Not Found


-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

look at the connection:
- you can use google chrome dev tools / network
- you can use CURL at the command line
	curl -v golang.org

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

https://en.wikipedia.org/wiki/List_of_HTTP_header_fields

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-