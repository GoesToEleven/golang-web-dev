google: http request specification
-------------------------------------
https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html
https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html
https://tools.ietf.org/html/rfc7230#section-2


google: http status codes
-------------------------------------
https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
https://www.ietf.org/rfc/rfc2616.txt


https://tools.ietf.org/html/rfc7230#section-2
(capitalization sometimes mine)
---------------------------------------------
					           Standards Track
RFC 7230           HTTP/1.1 Message Syntax and Routing         June 2014
---------------------------------------------

HTTP was created for the World Wide Web (WWW) architecture and has
evolved over time to support the scalability needs of a worldwide
hypertext system.  Much of that architecture is reflected in the
terminology and syntax productions used to define HTTP.

2.1.  Client/Server Messaging

HTTP is a stateless request/response protocol that operates by
exchanging messages across a reliable TRANSPORT- or
SESSION-layer "CONNECTION".


An HTTP "CLIENT" is a program that establishes a CONNECTION to a server
for the purpose of sending one or more HTTP requests.

An HTTP "SERVER" is a program that accepts CONNECTIONS in order to service HTTP requests by sending HTTP responses.

The terms "client" and "server" refer only to the roles that these
programs perform for a particular connection.  The same program might
act as a client on some connections and a server on others.  The term
"user agent" refers to any of the various client programs that
initiate a request, including (but not limited to) browsers, spiders
(web-based robots), command-line tools, custom applications, and
mobile apps.  The term "origin server" refers to the program that can
originate authoritative responses for a given target resource.  The
terms "sender" and "recipient" refer to any implementation that sends
or receives a given message, respectively.

HTTP relies upon the Uniform Resource Identifier (URI) standard
to indicate the target resource and relationships between resources.
Messages are passed in a format similar to that used by Internet mail
and the Multipurpose Internet Mail Extensions (MIME) (see Appendix A of
[RFC7231] for the differences between HTTP and MIME messages).

Most HTTP communication consists of a retrieval request (GET) for a
representation of some resource identified by a URI.  In the simplest
case, this might be accomplished via a single bidirectional
CONNECTION (===) between the user agent (UA) and the origin
server (O).

	request   >
UA ======================================= O
						   <   response

A client sends an HTTP REQUEST to a server in the form of a REQUEST
MESSAGE, beginning with a REQUEST-LINE that includes a method, URI,
and protocol version, followed by header fields
containing request modifiers, client information, and representation
metadata, an empty line to indicate the end of the
header section, and finally a message body containing the payload
body (if any, Section 3.3).

A server responds to a client's request by sending one or more HTTP
RESPONSE MESSAGES, each beginning with a STATUS LINE that includes
the protocol version, a success or error code, and textual reason
phrase, possibly followed by header fields containing
server information, resource metadata, and representation metadata,
an empty line to indicate the end of the header
section, and finally a message body containing the payload body (if
any, Section 3.3).

A connection might be used for multiple REQUEST/RESPONSE exchanges,
as defined in Section 6.3.

The following example illustrates a typical message exchange for a
GET request on the URI
"http://www.example.com/hello.txt":

Client request:

GET /hello.txt HTTP/1.1
User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7l zlib/1.2.3
Host: www.example.com
Accept-Language: en, mi


Server response:

HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Server: Apache
Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
ETag: "34aa387-d-1568eb00"
Accept-Ranges: bytes
Content-Length: 51
Vary: Accept-Encoding
Content-Type: text/plain

Hello World! My payload includes a trailing CRLF.

2.2.  Implementation Diversity

When considering the design of HTTP, it is easy to fall into a trap
of thinking that all user agents are general-purpose browsers and all
origin servers are large public websites.  That is not the case in
practice.  Common HTTP USER AGENTS include household appliances,
stereos, scales, firmware update scripts, command-line programs,
mobile apps, and communication devices in a multitude of shapes and
sizes.  Likewise, common HTTP ORIGIN SERVERS include home automation
units, configurable networking components, office machines,
autonomous robots, news feeds, traffic cameras, ad selectors, and
video-delivery platforms.

The term "USER AGENT" does not imply that there is a human user
directly interacting with the software agent at the time of a
request.  In many cases, a user agent is installed or configured to
run in the background and save its results for later inspection (or
save only a subset of those results that might be interesting or
erroneous).  Spiders, for example, are typically given a start URI
and configured to follow certain behavior while crawling the Web as a
hypertext graph.

...





























2.3.  Intermediaries

HTTP enables the use of intermediaries to satisfy requests through a
chain of connections.  There are three common forms of HTTP
intermediary: PROXY, GATEWAY, and TUNNEL.  In some cases, a single
intermediary might act as an origin server, proxy, gateway, or
tunnel, switching behavior based on the nature of each request.

	>             >             >             >
UA =========== A =========== B =========== C =========== O
		  <             <             <             <

The figure above shows three intermediaries (A, B, and C) between the
USER AGENT and ORIGIN SERVER.  A request or response message that
travels the whole chain will pass through four separate connections.
Some HTTP communication options might apply only to the connection
with the nearest, non-tunnel neighbor, only to the endpoints of the
chain, or to all connections along the chain.  Although the diagram
is linear, each participant might be engaged in multiple,
simultaneous communications.  For example, B might be receiving
requests from many clients other than A, and/or forwarding requests
to servers other than C, at the same time that it is handling A's
request.  Likewise, later requests might be sent through a different
path of connections, often based on dynamic configuration for load
balancing.

The terms "upstream" and "downstream" are used to describe
directional requirements in relation to the message flow: all
messages flow from UPSTREAM to DOWNSTREAM.  The terms "inbound" and
"outbound" are used to describe directional requirements in relation
to the request route: "INBOUND" means toward the origin server and
"OUTBOUND" means toward the user agent.

A "PROXY" is a message-forwarding agent that is selected by the
client, usually via local configuration rules, to receive requests
for some type(s) of absolute URI and attempt to satisfy those
requests via translation through the HTTP interface.  Some
translations are minimal, such as for proxy requests for "http" URIs,
whereas other requests might require translation to and from entirely
different application-level protocols.  Proxies are often used to
group an organization's HTTP requests through a common intermediary
for the sake of security, annotation services, or shared caching.
Some proxies are designed to apply transformations to selected
messages or payloads while they are being forwarded, as described in
Section 5.7.2.

A "GATEWAY" (a.k.a. "reverse proxy") is an intermediary that acts as
an origin server for the outbound connection but translates received
requests and forwards them inbound to another server or servers.
Gateways are often used to encapsulate legacy or untrusted
information services, to improve server performance through
"accelerator" caching, and to enable partitioning or load balancing
of HTTP services across multiple machines.

All HTTP requirements applicable to an origin server also apply to
the outbound communication of a gateway.  A gateway communicates with
inbound servers using any protocol that it desires, including private
extensions to HTTP that are outside the scope of this specification.
However, an HTTP-to-HTTP gateway that wishes to interoperate with
third-party HTTP servers ought to conform to user agent requirements
on the gateway's inbound connection.

A "TUNNEL" acts as a blind relay between two connections without
changing the messages.  Once active, a tunnel is not considered a
party to the HTTP communication, though the tunnel might have been
initiated by an HTTP request.  A tunnel ceases to exist when both
ends of the relayed connection are closed.  Tunnels are used to
extend a virtual connection through an intermediary, such as when
Transport Layer Security (TLS, [RFC5246]) is used to establish
confidential communication through a shared firewall proxy.

The above categories for intermediary only consider those acting as
participants in the HTTP communication.  There are also
intermediaries that can act on lower layers of the network protocol
stack, filtering or redirecting HTTP traffic without the knowledge or
permission of message senders.  Network intermediaries are
indistinguishable (at a protocol level) from a man-in-the-middle
attack, often introducing security flaws or interoperability problems
due to mistakenly violating HTTP semantics.

For example, an "interception proxy" (also commonly known
as a "transparent proxy" or "captive portal") differs from
an HTTP proxy because it is not selected by the client.  Instead, an
interception proxy filters or redirects outgoing TCP port 80 packets
(and occasionally other common port traffic).  Interception proxies
are commonly found on public network access points, as a means of
enforcing account subscription prior to allowing use of non-local
Internet services, and within corporate firewalls to enforce network
usage policies.

HTTP is defined as a stateless protocol, meaning that each request
message can be understood in isolation.  Many implementations depend
on HTTP's stateless design in order to reuse proxied connections or
dynamically load balance requests across multiple servers.  Hence, a
server MUST NOT assume that two requests on the same connection are
from the same user agent unless the connection is secured and
specific to that agent.  Some non-standard HTTP extensions (e.g.,
 have been known to violate this requirement, resulting in
security and interoperability problems.

2.4.  Caches

A "CACHE" is a local store of previous response messages and the
subsystem that controls its message storage, retrieval, and deletion.
A cache stores cacheable responses in order to reduce the response
time and network bandwidth consumption on future, equivalent
requests.  Any client or server may employ a cache, though a cache
cannot be used by a server while it is acting as a tunnel.

The effect of a cache is that the request/response chain is shortened
if one of the participants along the chain has a cached response
applicable to that request.  The following illustrates the resulting
chain if B has a cached copy of an earlier response from O (via C)
for a request that has not been cached by UA or A.

	   >             >
  UA =========== A =========== B - - - - - - C - - - - - - O
			 <             <

A response is "cacheable" if a cache is allowed to store a copy of
the response message for use in answering subsequent requests.  Even
when a response is cacheable, there might be additional constraints
placed by the client or by the origin server on when that cached
response can be used for a particular request.  HTTP requirements for
cache behavior and cacheable responses are defined in Section 2 of
[RFC7234].

There is a wide variety of architectures and configurations of caches
deployed across the World Wide Web and inside large organizations.
These include national hierarchies of proxy caches to save
transoceanic bandwidth, collaborative systems that broadcast or
multicast cache entries, archives of pre-fetched cache entries for
use in off-line or high-latency environments, and so on.