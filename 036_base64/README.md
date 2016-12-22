In terms of actual standards, there have been a few attempts to codify cookie behaviour but none thus far actually reflect the real world.

RFC 2109 was an attempt to codify and fix the original Netscape cookie_spec. In this standard many more special characters are disallowed, as it uses RFC 2616 tokens (a - is still allowed there), and only the value may be specified in a quoted-string with other characters. No browser ever implemented the limitations, the special handling of quoted strings and escaping, or the new features in this spec.
RFC 2965 was another go at it, tidying up 2109 and adding more features under a ‘version 2 cookies’ scheme. Nobody ever implemented any of that either. This spec has the same token-and-quoted-string limitations as the earlier version and it's just as much a load of nonsense.
RFC 6265 is an HTML5-era attempt to clear up the historical mess. It still doesn't match reality exactly but it's much better then the earlier attempts—it is at least a proper subset of what browsers support, not introducing any syntax that is supposed to work but doesn't (like the previous quoted-string).
In 6265 the cookie name is still specified as an RFC 2616 token, which means you can pick from the alphanums plus:

!#$%&'*+-.^_`|~

In the cookie value it formally bans the (filtered by browsers) control characters and (inconsistently-implemented) non-ASCII characters. It retains cookie_spec's prohibition on space, comma and semicolon, plus for compatibility with any poor idiots who actually implemented the earlier RFCs it also banned backslash and quotes, other than quotes wrapping the whole value (but in that case the quotes are still considered part of the value, not an encoding scheme). So that leaves you with the alphanums plus:

!#$%&'()*+-./:<=>?@[]^_`{|}~
In the real world we are still using the original-and-worst Netscape cookie_spec, so code that consumes cookies should be prepared to encounter pretty much anything, but for code that produces cookies it is advisable to stick with the subset in RFC 6265.

http://stackoverflow.com/questions/1969232/allowed-characters-in-cookies