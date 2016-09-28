Create a new type "snoop" with an underlying type of int.

Attach a method to type "snoop" so that any value of type "snoop" implements the http.Handler interface.

In func main, create a variable with the identifier "dog" which is of type "snoop"

Create a new ServeMux.

For the route "/", use "dog" as the http.Handler argument.

When someone goes to the route "/", set the header "Content-Type" to "text/html; charset=utf-8"

Also use io.WriteString to write the "req.URL.Path" to the response.

ListenAndServe on port ":8080" using your ServeMux.