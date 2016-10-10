# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go
	fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.