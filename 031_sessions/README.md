# Session

This is how we create state:

We will store a Universally unique identifier [(UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier) in the cookie. 

We will associate each user with a UUID.

This will allow us to identify every user visiting our website.

## Security 

There are three factors which contribute to the security of a session created using a cookie and a UUID:

1. Uniqueness of the ID
1. Encryption in transit with HTTPS
1. Encryption in storage with HMAC

A UUID is very unique.

[Wikipedia says this about UUIDs:](https://en.wikipedia.org/wiki/Universally_unique_identifier) " ... only after generating 1 billion UUIDs every second for the next 100 years, the probability of creating just one duplicate would be about 50%."

A UUID cannot be intercepted in transit if we are using HTTPS.

A UUID is secure in stoage if we encrypt the UUID with a Hash-based Message Authentication Code [(HMAC)](https://en.wikipedia.org/wiki/Hash-based_message_authentication_code).

We will look at HTTPS and HMAC in the next two sections.





[StackOverflow #1](http://security.stackexchange.com/questions/30707/demystifying-web-authentication-stateless-session-cookies)

[StackOverflow #2](http://security.stackexchange.com/questions/7398/secure-session-cookies)

[StackOverflow #3](http://stackoverflow.com/questions/3240246/signed-session-cookies-a-good-idea)

[StackOverflow #4](http://stackoverflow.com/questions/5459682/how-to-protect-session-id-and-cookies-without-using-ssl-https)

# Hash Based Message Authentication (HMAC)

[Wikipedia HMAC](https://en.wikipedia.org/wiki/Hash-based_message_authentication_code)

[Wikipedia MAC](https://en.wikipedia.org/wiki/Message_authentication_code)

[Wikipedia Replay Attack](https://en.wikipedia.org/wiki/Replay_attack)

[Stack Overflow HMAC](http://security.stackexchange.com/questions/20129/how-and-when-do-i-use-hmac/20301)



