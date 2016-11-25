We shouldn't store unencrypted passwords.

Use the golang.org/x/crypto/bcrypt package to:
- encrypt your password
- compare stored password with entered password on login

Do not worry about abstracting code. You may notice redundant code. This is fine.