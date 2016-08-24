# Setup For App Engine

Make sure your environment is configured for App Engine.

If you have not worked with App Engine yet, please see "Hello, World!" for Go on App Engine

https://cloud.google.com/appengine/docs/go/

# Setup for Google Cloud Storage

You need to install Google Cloud Storage SDK.

https://cloud.google.com/sdk/

# Create A Project on Google Cloud Platform (App Engine)

If you have not already, make sure you have a project on Google Cloud Platform 

https://console.cloud.google.com/project

# Create A Default Cloud Storage Bucket

At the time I wrote this ...

You find this by clicking the top-left hamburger menu

... then choosing ...

COMPUTER / APP ENGINE / SETTINGS / APPLICATION SETINGS 	

... or by following this link ...

https://console.cloud.google.com/appengine/settings

... and then clicking on the CREATE button at the bottom ...

Create
Default Cloud Storage Bucket
A free 5GB Cloud Storage bucket for App Engine applications, doesn't require billing enabled.

# Download These Packages

go get -u golang.org/x/oauth2

go get -u google.golang.org/cloud/storage

go get -u google.golang.org/appengine/...

# Configure Your App.yaml

Make sure it looks like this:

```go
application: <your-app-id-here>
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
script: _go_app
```

# FYI - Reading Documentation

When you look up documentation on godoc.org, go to the parent package, then look at "Directories"

## Example

For instance, if I wanted to see the documentation on this package:

google.golang.org/appengine/file

I would go to this package

google.golang.org/appengine

Then click on "Directories" (or scroll to the bottom), find the "file" package, then click that link.


## Example

For instance, if I wanted to see the documentation on this package:

google.golang.org/cloud/storage

I would go to this package

google.golang.org/cloud

Then click on "Directories" (or scroll to the bottom), find the "storage" package, then click that link.

# To Remember

## Files Implement The Reader Interface

https://godoc.org/os#pkg-index

Type `*os.File` has the method `func (f *File) Read(b []byte) (n int, err error)` 

This is the reader interface - type `io.Reader`:

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

Anything with a `Read(p []byte) (n int, err error)` method impicitly implements the reader interface and is of type io.Reader.

## Multipart Files Implement The Reader Interface

In our first example, we will use `io.Copy` which is defined like this:

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

We will be able to pass a multi-part file into Copy because a multi-part file implements the reader interface.

```go
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
```

https://godoc.org/mime/multipart#File

# The Process For STORING in GCS

Reference the index here when going through the steps below:

https://godoc.org/google.golang.org/cloud/storage#pkg-index

## Terminology

### Bucket

Buckets are the basic containers that hold your data. Everything that you store in Google Cloud Storage must be contained in a bucket. You can use buckets to organize your data and control access to your data, but unlike directories and folders, you cannot nest buckets. There is a per-project rate limit on bucket create/delete operations of approximately 1 operation every 2 seconds; however there is no rate limit on object create/delete operations. Therefore, we recommend that you design your storage applications to favor intensive object operations and relatively few buckets operations.

When you create a bucket, you specify a storage class that best matches the type of data the bucket will contain. For example, for frequently accessed data, you can use the Standard Storage class. For infrequently accessed data, you can use the Cloud Storage Nearline class.

## Object

Objects are the individual pieces of data that you store in Google Cloud Storage. A single object can be up to 5 TB in size. Objects have two components: object data and object metadata. The object data component is usually a file that you want to store in Google Cloud Storage. The object metadata component is a collection of name-value pairs that describe various object qualities. There is no limit on the number of objects that you can create in a bucket.

Objects are immutable, which means that an uploaded object cannot change throughout its storage lifetime. An object's storage lifetime is the time between successful object creation (upload) and successful object deletion. In practice, this means that you cannot make incremental changes to objects, such as append operations or truncate operations. However, it is possible to overwrite objects that are stored in Google Cloud Storage, and doing so happens atomically — until the new upload completes the old version of the object will be served to readers, and after the upload completes the new version of the object will be served to readers. So a single overwrite operation simply marks the end of one immutable object's lifetime and the beginning of a new immutable object's lifetime.

There is no limit to how quickly you can create or update different objects in a bucket. However, a single particular object can only be updated or overwritten up to once per second. For example, if you have an object bar in bucket foo, then you should only upload a new copy of foo/bar about once per second. Updating the same object faster than once per second may result in 503 Service Unavailable errors.

An object name is just metadata to Google Cloud Storage. Object names can contain any combination of Unicode characters (UTF-8 encoded) less than 1024 bytes in length. A common character to include in file names is a slash (/). By using slashes in an object name, you can make objects appear as though they're stored in a hierarchical structure. For example, you could name one object /europe/france/paris.jpg and another object /europe/france/cannes.jpg. When you list these objects, they appear to be in a hierarchical directory structure based on location; however, Google Cloud Storage sees the objects as independent objects with no hierarchical relationship whatsoever.

Source:
https://cloud.google.com/storage/docs/key-terms

## Code Sample - Storing An Object

In it's most basic form, our code for storing an object will look like this:

```go
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, someFile)
	return writer.Close()
```
 
Pay particular attention to these lines:

```go
	...
	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	...
	io.Copy(writer, someFile)
	...
}
```

We will create a google cloud client.

We specify the bucket name.

We specify the object name.

We call `NewWriter(ctx)`.

### Interface Polymorphism Coolness

In the example code above, we have `io.Copy(writer, someFile)`.

`io.Copy` actually takes a writer and a reader.

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

We can put any type that implements the Reader interface in as an argument for the `src Reader` parameter.

### NewWriter

NewWriter returns a storage Writer that writes to the GCS object associated with this ObjectHandle.

A new object will be created if an object with this name already exists. Otherwise any previous object with the same name will be replaced. The object will not be available (and any previous object will remain) until Close has been called.

Attributes can be set on the object by modifying the returned Writer's ObjectAttrs field before the first call to Write. If no ContentType attribute is specified, the content type will be automatically sniffed using net/http.DetectContentType.

It is the caller's responsibility to call Close when writing is done.

Source:
https://godoc.org/google.golang.org/cloud/storage#ObjectHandle.NewWriter

### Look at the documentation to reinforce this

https://godoc.org/google.golang.org/cloud/storage#pkg-index

What other possibilities are there?

When we looked at the documentation for `NewWriter(ctx)` it said that it is the caller's responsibility to call Close when writing is done.

In the documentation, there are several `Close` methods attached to different types.

Which `Close` method, attached to which type, do we want?

## Code Sample - Retrieving An Object

In it's most basic form, our code for storing an object will look like this:

```go
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(gcsBucket).Object(name).NewReader(ctx)
```

# Troubleshooting

If any of the examples don't work, try clearing cookies and then running the example again. To keep code samples specific and short, I kept logic to a minimum.


