# Third-party ServeMux

You can search [godoc.org](https://godoc.org/) for third-party packages.

Here is [a good third-party ServeMux](https://godoc.org/github.com/julienschmidt/httprouter) that allows easy access to methods for routing & path parameters.

# [julienschmidt/httprouter](https://godoc.org/github.com/julienschmidt/httprouter)

## Match method & path

The router matches incoming requests by the request method and the path.
 
 ``` Go
 func main() {
     router := httprouter.New()
     router.GET("/apply", apply)
     router.POST("/apply", applyProcess) 
     http.ListenAndServe(":8080", router)
 }
 ```

## Named path parameters

The registered path, against which the router matches incoming requests, can also contain parameters. Parameters are dynamic path segments. They match anything until the next '/' or the path end.

``` Go
func main() {
    router := httprouter.New()
    router.GET("/blog/:category/:article", blog)
    http.ListenAndServe(":8080", router)
}
```

```
Requests:
 /blog/go/request-routers            match: category="go", article="request-routers"
 /blog/go/request-routers/           no match, but the router would redirect
 /blog/go/                           no match
 /blog/go/request-routers/comments   no match
```

## Catch-all path parameters

Catch-all parameters match anything until the path end, including the directory index (the '/' before the catch-all). Since they match anything until the end, catch-all parameters must always be the final path element.

``` Go
func main() {
    router := httprouter.New()
    router.GET("/files/*filepath", loadFile)
    http.ListenAndServe(":8080", router)
}
```

```
Requests:
 /files/                             match: filepath="/"
 /files/LICENSE                      match: filepath="/LICENSE"
 /files/templates/article.html       match: filepath="/templates/article.html"
 /files                              no match
```

## Using path parameters

The value of parameters is saved as a []Param


``` Go
type Param struct {
    Key   string
    Value string
}
```

The slice is passed to the Handle func as a third parameter. 

``` Go
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("user"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:user", Hello)

    http.ListenAndServe(":8080", router)
}
```

Retrieve the value of a parameter:

``` Go
user := ps.ByName("user") // defined by :user or *user

```

## Performance

Julien Schmidt's router is [nicely performant](https://github.com/julienschmidt/go-http-routing-benchmark#static-routes)

