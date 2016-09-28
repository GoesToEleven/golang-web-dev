Modify the code from the previous problem so that the path of a url is split on the "/" and then the response writes back the parts of the path separated by dashes. For instance:

http://localhost:8080/todd/mcleod/dad

Prints out

- todd - mcleod - dad

Use:
req.URL.Path
strings.Split
strings.Join