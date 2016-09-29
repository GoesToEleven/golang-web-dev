# TECHNIQUES WE WILL LEARN

- concatenate
- CLI pipeline - output to a file with > 

# CODE WE WILL USE FROM THE STANDARD LIBRARY

## [os.Create](https://godoc.org/os#Create)
This allows us to create a file.
``` Go
func Create(name string) (*File, error)
```

*** 

# DEFER
## defer keyword

***

# READING & WRITING

## [io.Copy](https://godoc.org/io#Copy)
``` Go
func Copy(dst Writer, src Reader) (written int64, err error)
```

## [strings.NewReader](https://godoc.org/strings#NewReader)
``` Go
func NewReader(s string) *Reader
```