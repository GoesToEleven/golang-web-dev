- concatenate
- pipeline output to a file with > 

- [os.Create](https://godoc.org/os#Create)
- func Create(name string) (*File, error)

- defer

- [io.Copy](https://godoc.org/io#Copy)
- func Copy(dst Writer, src Reader) (written int64, err error)

- [strings.NewReader](https://godoc.org/strings#NewReader)
- func NewReader(s string) *Reader
