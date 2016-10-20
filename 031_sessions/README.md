# Session

This is how we create state:

We will store a unique ID in the cookie. 

On our server, we will associate each user with a unique ID.

This will allow us to identify every user visiting our website.

## Security 

There are two factors which contribute to the security of a session created using a cookie and a unique ID:

1. Uniqueness of the ID
1. Encryption in transit with HTTPS

You can use any unique ID that you would like: a  Universally unique identifier [(UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier) or a database key. If you're using a database key, make sure it's not the key to your user but to a separate session table.

A UUID is very unique. [Wikipedia says this about UUIDs:](https://en.wikipedia.org/wiki/Universally_unique_identifier) " ... only after generating 1 billion UUIDs every second for the next 100 years, the probability of creating just one duplicate would be about 50%."

A UUID cannot be intercepted in transit if we are using HTTPS.

We will look at HTTPS in the next section.