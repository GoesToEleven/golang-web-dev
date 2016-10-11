# ListenAndServe on port 8080 of localhost

For the default route "/"
Have a func called "foo" 
which writes to the response "foo ran"

For the route "/dog/"
Have a func called "dog" 
which parses a template called "dog.gohtml"
and writes to the response "<h1>This is from dog</h1>"
and also shows a picture of a dog when the template is executed.

Use "http.ServeFile"
to serve the file "dog.jpeg"