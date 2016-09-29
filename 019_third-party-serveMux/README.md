# Third-party ServeMux

You can search [godoc.org](https://godoc.org/) for third-party packages.

Here is [a good third-party ServeMux]() that allows easy access to methods for routing & path parameters.

## [julienschmidt/httprouter](https://godoc.org/github.com/julienschmidt/httprouter)

The router matches incoming requests by the request method and the path.
 
 ``` Go
 func main() {
     router := httprouter.New()
     router.GET("/apply", apply)
     router.POST("/apply", applyProcess) 
     log.Fatal(http.ListenAndServe(":8080", router))
 }
 ```

The registered path, against which the router matches incoming requests, can also contain parameters. Parameters are dynamic path segments. They match anything until the next '/' or the path end.

``` Go
func main() {
    router := httprouter.New()
    router.GET("/blog/:category/:post", blog)
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

```
Requests:
 /blog/go/request-routers            match: category="go", post="request-routers"
 /blog/go/request-routers/           no match, but the router would redirect
 /blog/go/                           no match
 /blog/go/request-routers/comments   no match
```



