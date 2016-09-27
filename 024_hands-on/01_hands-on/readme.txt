ListenAndServe on port 8080 of localhost

For the default route "/"
Have a HandlerFunc called "foo" which writes to the response "foo ran"

For the default route "/uptown/"
Have a HandlerFunc called "dog" which parses a template called "fido.gohtml"
and writes to the response "<h1>This is from dog</h1>"
and also shows a picture of a dog when the template is executed.

For the default route "/fido/"
Have a HandlerFunc called "chien" which serves the file "dog.jpeg"