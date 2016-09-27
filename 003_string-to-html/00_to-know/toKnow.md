- concatenate
- pipeline output to a file with > 

- [os.Create](https://godoc.org/os#Create)
``` Go
func Create(name string) (*File, error)
```

- defer

- [io.Copy](https://godoc.org/io#Copy)
``` Go
func Copy(dst Writer, src Reader) (written int64, err error)
```
- [strings.NewReader](https://godoc.org/strings#NewReader)
``` Go
func NewReader(s string) *Reader
```