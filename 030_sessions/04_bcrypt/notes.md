You should never store a password without encrypting it.

We will use ```bcrypt``` to encrypt our passwords.

# Step 1: 

You will need to go get this package:

```
go get golang.org/x/crypto/bcrypt
```

# Step 2:
Change your user struct's password field to the data type []byte

# Step 3:
Use bcrypt to encrypt your password before storing it.
