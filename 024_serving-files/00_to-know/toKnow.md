# DIFFERENT WAYS TO SERVE

********************************

# io.COPY

[os.Open](https://godoc.org/os#Open)
``` go
func Open(name string) (*File, error)
```

[os.File.Read](https://godoc.org/os#File.Read)
``` go
func (f *File) Read(b []byte) (n int, err error)
```

[io.Copy](https://godoc.org/io#Copy)
``` go
func Copy(dst Writer, src Reader) (written int64, err error)
```

********************************

# SERVECONTENT

[http.ServeContent](https://godoc.org/net/http#ServeContent)
``` go
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
```

********************************

# SERVEFILE

[http.ServeFile](https://godoc.org/net/http#ServeFile)
``` go
func ServeFile(w ResponseWriter, r *Request, name string)
```
********************************

# FILESERVER & STRIPPREFIX

[http.FileServer](https://godoc.org/net/http#FileServer)
``` go
func FileServer(root FileSystem) Handler
```

[http.StripPrefix](https://godoc.org/net/http#StripPrefix)
``` go
func StripPrefix(prefix string, h Handler) Handler
```

********************************
