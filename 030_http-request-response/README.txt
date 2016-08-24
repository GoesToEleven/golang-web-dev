Request and response messages are similar. Both messages consist of:

- an request/response line
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

https://en.wikipedia.org/wiki/List_of_HTTP_header_fields

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

SOURCES:
https://www.jmarshall.com/easy/http/
https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html