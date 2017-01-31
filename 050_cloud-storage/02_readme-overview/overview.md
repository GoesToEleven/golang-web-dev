# Put A File

We put a file with code like this:

```go
func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)

	io.Copy(writer, rdr)
	// check for errors on io.Copy in production code!
	return writer.Close()
}
```

You can learn about the different types and methods in google/cloud by going to godoc.org and searching for `google/cloud` which will give you this link: [https://godoc.org/google.golang.org/cloud/storage](https://godoc.org/google.golang.org/cloud/storage)

My read through on the code would be like this:

In the `storage` package for Google cloud, you call `func newClient` which gives you a pointer to a client: `*storage.Client`.

Once you have a pointer to a client, there are two methods which you can call:

```go
func (c *Client) Bucket(name string) *BucketHandle
```

```go
func (c *Client) Close() error
```

Since there is a `Close()` method, this is a big hint that we should probably close a client. We can do this effectively with defer, as you see in the code sample above: `defer client.Close()`

To continue finding functionality in Google Cloud Storage, we will then also probably need to call `func (c *Client) Bucket(name string) *BucketHandle`. This makes sense as we will need to specify which bucket we want to access either to **put** or **get** a file.

So if we call `Bucket()` we are given a new type which is a pointer to a bucket handle: `*storage.BucketHandle`.

With a `*storage.BucketHandle`, we can see in the `storage` package documentation index that there are now several more methods available to us:
 
 ```go
 func (c *BucketHandle) ACL() *ACLHandle
 func (b *BucketHandle) Attrs(ctx context.Context) (*BucketAttrs, error)
 func (c *BucketHandle) DefaultObjectACL() *ACLHandle
 func (b *BucketHandle) List(ctx context.Context, q *Query) (*ObjectList, error)
 func (b *BucketHandle) Object(name string) *ObjectHandle
 ```
 
## ACL
 
This will let us control our **Access Control List**. Basically these are settings which we can set on our **bucket** to control who can access the **bucket** and what they can do. This is known as *scope* and *permissions*.
  
**Scope** defines who the permission applies to (for example, a specific user or group of users). Scopes are sometimes referred to as *grantees.* 

**Permissions** define the actions that can be performed against a bucket (for example, read or write).
   
We will also, later, be able to set ACL's for objects (files) which we store in Google Cloud Storage (GCS).
  
More on this later.
 
## Attrs
 
This will give us the attributes for a **bucket**. You can see the many different attributes at this link: [https://godoc.org/google.golang.org/cloud/storage#BucketAttrs](https://godoc.org/google.golang.org/cloud/storage#BucketAttrs) For convenience, I'm also listing them here:
 
```go
 type BucketAttrs struct {
     // Name is the name of the bucket.
     Name string
 
     // ACL is the list of access control rules on the bucket.
     ACL []ACLRule
 
     // DefaultObjectACL is the list of access controls to
     // apply to new objects when no object ACL is provided.
     DefaultObjectACL []ACLRule
 
     // Location is the location of the bucket. It defaults to "US".
     Location string
 
     // MetaGeneration is the metadata generation of the bucket.
     MetaGeneration int64
 
     // StorageClass is the storage class of the bucket. This defines
     // how objects in the bucket are stored and determines the SLA
     // and the cost of storage. Typical values are "STANDARD" and
     // "DURABLE_REDUCED_AVAILABILITY". Defaults to "STANDARD".
     StorageClass string
 
     // Created is the creation time of the bucket.
     Created time.Time
 }
 ```
 
We will also, later, be able to see Attrs for objects (files) which we store in Google Cloud Storage (GCS).
  
More on this later.
  
## DefaultObjectACL
  
This let's you set a default ACL which will be applied to newly created objects in this bucket that do not have a defined ACL.

## List

List lists objects from the bucket. You can specify a query to filter the results. If q is nil, no filtering is applied.

This is what we will use to query a bucket and have results returned.

```go
func (b *BucketHandle) List(ctx context.Context, q *Query) (*ObjectList, error)
```

## Object

This is perhaps the most commonly used method when working with a bucket.

Remember what we've done so far: We (1) got a Google Cloud Storage client, and then we (2) said that we wanted to work with a specific bucket, and now (3) we are going to say that we want to work with a specific object.

The code to do all of that is in our initial code sample up above. The excerpt of code to which I'm referring looks like this:


```go
 client.Bucket(gcsBucket).Object(name)
```

As we are learning how to **put** an object here, we will follow this thread of logic.

The `Object` method returns a pointer to an ObjectHandle `*storage.ObjectHandle`. You can see the func's signature here again:

```go
func (b *BucketHandle) Object(name string) *ObjectHandle
```

With a `*ObjectHandle` we once again have several methods available to us:
 
 ```go
 func (o *ObjectHandle) ACL() *ACLHandle
 func (o *ObjectHandle) Attrs(ctx context.Context) (*ObjectAttrs, error)
 func (o *ObjectHandle) CopyTo(ctx context.Context, dst *ObjectHandle, attrs *ObjectAttrs) (*ObjectAttrs, error)
 func (o *ObjectHandle) Delete(ctx context.Context) error
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64) (*Reader, error)
 func (o *ObjectHandle) NewReader(ctx context.Context) (*Reader, error)
 func (o *ObjectHandle) NewWriter(ctx context.Context) *Writer
 func (o *ObjectHandle) Update(ctx context.Context, attrs ObjectAttrs) (*ObjectAttrs, error)
 func (o *ObjectHandle) WithConditions(conds ...Condition) *ObjectHandle
 ```

You can see that, for an **object** (and not the bucket), we can set the **ACL** (Access Control List) for an **object**, we can also see the **Attrs** for an object, we can use **CopyTo** to copy one object to another object, we can **Delete** an object, we can read an object with **NewReader** and we can write to an object with **NewWriter**.

We can also **Update** the attributes ( `ObjectAttrs` ) of an object. Please note that you *cannot* alter the actual binary of the object - there is no prepending or appending something to a video file, there is no adding binary to a picture, none of that. If you want to *change* the actual object, you need to *replace* the object with a new version. However, if you only want to *update* the *attributes* of an object, well then, you can use the **Update** method.
  
Regarding the last method, **WithConditions**, don't worry about this one right now. You can learn about it if, and when, you need it.

So to **put** an object, we're going to call **NewWriter** because, what we want to do, is write to this one particular object in this one particular bucket in Google Cloud Storage. We can see this from our initial code samply up above:

```go
writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
```

Now that we have a writer, we can use `io.Copy` to copy from something that implements the reader interface into our new object on GCS:

```go
	io.Copy(writer, rdr)
	// check for errors on io.Copy in production code!
	return writer.Close()
```

And, as you can see above, we should close our writer. We know this because, when we call `NewWriter` we are returned a pointer to a google cloud storage writer:

```go
func (o *ObjectHandle) NewWriter(ctx context.Context) *Writer
```

And with a `*storage.Writer`, we have these methods:
 
 ```go
  func (w *Writer) Attrs() *ObjectAttrs
  func (w *Writer) Close() error
  func (w *Writer) CloseWithError(err error) error
  func (w *Writer) Write(p []byte) (n int, err error)
 ```

You can see in the above methods that the type `storage.Writer` implements the `io.Writer` interface. Here is that interface from package `io`:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Since a `*storage.Writer` has a method with a signature `Write(p []byte) (n int, err error)` that matches the required method to implement the `io.Writer` interface, then the type `storage.Writer` implicitly impelements the writer interface.

That means we can use a `sotrage.Writer` in any func or method which asks for a `Writer`.

You can see this with `io.Copy`:

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

`io.Copy` is asking for a `Writer` so, since `storage.Writer` implements the writer interface, we can pass `io.Copy` a `storage.Writer`.

Ok. So that is a little bit about how you put an object in GCS, and a lot about reading documentation and understanding Go code.

Now onto how we get an object.

# Get

This is the code to get a file from Google Cloud Storage:

```go
func getFile(ctx context.Context, name string) (io.ReadCloser, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Bucket(gcsBucket).Object(name).NewReader(ctx)
}
```

The process is very similar to how we *put* an object.

We get a storage **client**, and then we specify which **bucket** we want to access, and then we specify which **object** (file) we want to access, and then we call the method `NewReader` which gives us back a pointer to type Reader from package storage `*storage.Reader` and with a `*storage.Reader` we then have all of these methods available to us:

```go
func (r *Reader) Close() error
func (r *Reader) ContentType() string
func (r *Reader) Read(p []byte) (int, error)
func (r *Reader) Remain() int64
func (r *Reader) Size() int64
```

So you can see we are going to need to *close* our reader. Why? Well, just imagine what would happen to your computer if you just kept opening files after files after files to read them yourself *and you never closed any of those files*. Eventually, obviously, your computer is going to freak out which is another way of saying that it will run out of resources, like memory, and then just crash and die a painful death. So, to avoid all of that, close your files if you open them.

You will also notice that there is this method:

```go
func (r *Reader) Read(p []byte) (int, error)
```

What that tells us is that the type `storage.Reader` (and please notice that I keep prefixing the type with the package from which it comes - this is something that a lot of people need to be shown, the notation is `package.Type`, you see, just like here, `storage.Reader` we're talking about the type `Reader` from the `storage` package, ok). Ok. Back to what I was saying. What that tells us is that the type `storage.Reader` implements the reader interface from package io:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

You see, package `io` has a type `Reader` which is an interface. To implement that interface you have to have this method `Read(p []byte) (n int, err error)`. If there is a type that has that method, then it implements the `Reader` interface. And any func or method which asks for a value of type `Reader` as an argument can now take **any** type which implements the `Reader` interface.
 
Again, all interfaces are implicitly implemented in Go. No need to explicitly say that something implements an interface. It automatically happens implicitly.

So the type which we just got, `storage.Reader`, implements the reader interface. That means we can now use things like:

**io.Copy**

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

**bufio.NewReader**
```go
type Reader
func NewReader(rd io.Reader) *Reader
func NewReaderSize(rd io.Reader, size int) *Reader
func (b *Reader) Buffered() int
func (b *Reader) Discard(n int) (discarded int, err error)
func (b *Reader) Peek(n int) ([]byte, error)
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) ReadByte() (c byte, err error)
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
func (b *Reader) ReadString(delim byte) (line string, err error)
func (b *Reader) Reset(r io.Reader)
func (b *Reader) UnreadByte() error
func (b *Reader) UnreadRune() error
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
```

**bufio.NewScanner**
```go
type Scanner
func NewScanner(r io.Reader) *Scanner
func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string
```

**util.ReadAll**
```go
func ReadAll(r io.Reader) ([]byte, error)
```

**csv.NewReader**
```go
func NewReader(r io.Reader) *Reader
func (r *Reader) Read() (record []string, err error)
func (r *Reader) ReadAll() (records [][]string, err error)
```

**json.NewDecoder**
```go
type Decoder
func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Buffered() io.Reader
func (dec *Decoder) Decode(v interface{}) error
func (dec *Decoder) More() bool
func (dec *Decoder) Token() (Token, error)
func (dec *Decoder) UseNumber()
```

**fmt.Fscan**
```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
```

# Attrs.MediaLink

The next thing we'll look at is a media-link.

Please notice that the sequence of this document matches the sequence of the code samples in the folders which follow.

A mediaLink gives us a link which, when clicked, will allow an object (file) *to be downloaded*. You can read more about it [here](https://cloud.google.com/storage/docs/json_api/v1/objects)

Using this link could allow for some cost savings - your user's download something straight from GCS instead of going through you servers. If they went through your serves, that would require additional processing on the part of your application, for which you would get billed.

Ok. So to get a medialink, we use this code:

```go
func getFileLink(ctx context.Context, name string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	attrs, err := client.Bucket(gcsBucket).Object(name).Attrs(ctx)
	if err != nil {
		return "", err
	}
	return attrs.MediaLink, nil

```

My read through of the code would go like this:

We call `storage.NewClient` which returns to use a `*storage.Client` (a pointer to a storage client). 

With a `*storage.Client`, we can then call the `Bucket` method which is attached to a `*storage.Client`. 

That will return a `*storage.BucketHandle` (a pointer to a storage `BucketHandle`) to us.  

With a `*storage.BucketHandle`, we can then call the `Object` method which is attached to a `*storage.BucketHandle`.

That will return a `*storage.ObjectHandle` (a pointer to a storage `ObjectHandle`) to us.  

With a `*storage.ObjectHandle`, we can then call the `Attrs` method which is attached to a `*storage.ObjectHandle`.

That will return a `*storage.ObjectAttrs` (a pointer to a storage `ObjectAttrs`) to us.  

A `*storage.ObjectAttrs` is a memory address referencing a value which is a struct having these fields:

```go
type ObjectAttrs struct {
    // Bucket is the name of the bucket containing this GCS object.
    // This field is read-only.
    Bucket string

    // Name is the name of the object within the bucket.
    // This field is read-only.
    Name string

    // ContentType is the MIME type of the object's content.
    ContentType string

    // ContentLanguage is the content language of the object's content.
    ContentLanguage string

    // CacheControl is the Cache-Control header to be sent in the response
    // headers when serving the object data.
    CacheControl string

    // ACL is the list of access control rules for the object.
    ACL []ACLRule

    // Owner is the owner of the object. This field is read-only.
    //
    // If non-zero, it is in the form of "user-<userId>".
    Owner string

    // Size is the length of the object's content. This field is read-only.
    Size int64

    // ContentEncoding is the encoding of the object's content.
    ContentEncoding string

    // ContentDisposition is the optional Content-Disposition header of the object
    // sent in the response headers.
    ContentDisposition string

    // MD5 is the MD5 hash of the object's content. This field is read-only.
    MD5 []byte

    // CRC32C is the CRC32 checksum of the object's content using
    // the Castagnoli93 polynomial. This field is read-only.
    CRC32C uint32

    // MediaLink is an URL to the object's content. This field is read-only.
    MediaLink string

    // Metadata represents user-provided metadata, in key/value pairs.
    // It can be nil if no metadata is provided.
    Metadata map[string]string

    // Generation is the generation number of the object's content.
    // This field is read-only.
    Generation int64

    // MetaGeneration is the version of the metadata for this
    // object at this generation. This field is used for preconditions
    // and for detecting changes in metadata. A metageneration number
    // is only meaningful in the context of a particular generation
    // of a particular object. This field is read-only.
    MetaGeneration int64

    // StorageClass is the storage class of the bucket.
    // This value defines how objects in the bucket are stored and
    // determines the SLA and the cost of storage. Typical values are
    // "STANDARD" and "DURABLE_REDUCED_AVAILABILITY".
    // It defaults to "STANDARD". This field is read-only.
    StorageClass string

    // Created is the time the object was created. This field is read-only.
    Created time.Time

    // Deleted is the time the object was deleted.
    // If not deleted, it is the zero value. This field is read-only.
    Deleted time.Time

    // Updated is the creation or modification time of the object.
    // For buckets with versioning enabled, changing an object's
    // metadata does not change this property. This field is read-only.
    Updated time.Time
}
```

One of those fields is a `MediaLink` which is of type string and is what we return in our code sample at the beginning of this section on getting a medialink.

# Displaying An Image

The next thing which was interesting to me about GCS was getting an object which is an image from GCS and then displaying it on a webpage.

We saw how to download an object in the previous example.

Now let's take a look at how to get an object which is an image and then show it on a webpage.

The easiest way to understand this is to (1) upload an image to GCS, (2) go to your [google cloud console](), (3) navigate to your project, then to *storage* (at the time of this writing, use the top left "hamburger" menu), then to the bucket with your image, and then (4) click **public link** to allow the object to be shared publicly.

Once you have clicked **public link** you will be taken to a URL like this:

```html
https://storage.googleapis.com/learning-1130.appspot.com/3d79b3c8f88125cdc1fa1e0b82953460508f79e5.jpeg
```

This URL works for any image.

Add this prefix `https://storage.googleapis.com/learning-1130.appspot.com/` to the name of any object, and then you have a public link for that object.

# List Files

With the *list-files* example, the code we are using was originally from the [Google GCS Example](https://cloud.google.com/appengine/docs/go/googlecloudstorageclient/using-cloud-storage) and on [GitHub](https://github.com/GoogleCloudPlatform/gcloud-golang/tree/master/examples/storage/appengine). 

I have modified this code to focus upon one aspect of GCS at a time.

A nice aspect of this code is that a custom struct called `demo` is created, and then methods are attached to that struct.

```go
type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}
```

And here is an example of a method attached to `demo`:
 
```go
func (d *demo) listFiles() {
	io.WriteString(d.res, "\nRETRIEVING FILE NAMES\n")

	client, err := storage.NewClient(d.ctx)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}
	defer client.Close()

	objs, err := client.Bucket(gcsBucket).List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.res, obj.Name+"\n")
	}
}
```

You can see that the procedure here is very similar to what we saw before: Get a client, get a bucket.

Instead of getting an object, however, we're calling the `List` method which returns a `*storage.ObjectList`.

This is an `ObjectList`:

```go
type ObjectList struct {
    // Results represent a list of object results.
    Results []*ObjectAttrs

    // Next is the continuation query to retrieve more
    // results with the same filtering criteria. If there
    // are no more results to retrieve, it is nil.
    Next *Query

    // Prefixes represents prefixes of objects
    // matching-but-not-listed up to and including
    // the requested delimiter.
    Prefixes []string
}
```

We have **one** pointer to an `ObjectList` returned. *This is a list of objects.* It is **one** list. It is not many objects. We are not returned many objects. We are returned **one** list.

This ObjectList has three fields. The `Results` field is a slice which, like all slices, contains zero-to-many values. The values in this slice are of type `*storage.ObjectAttrs`. We've already looked at both [`storage.ObjectAttrs`](https://godoc.org/google.golang.org/cloud/storage#ObjectAttrs) and [`storage.BucketAttrs`](https://godoc.org/google.golang.org/cloud/storage#BucketAttrs). Remember that both of these types are structs which contain fields which are attributes of either an object or a bucket, respectively.

So with the `List` method, we now have access to all of the attributes of all of the objects in a specific bucket.

This means we can now get things like the `Name` of every object in a bucket, or the `MediaLink` for every object in a bucket, or any other variety of Attrs for each object in a bucket.

In the first example `list-files` we print out the name of each object.

In the second example `object-attributes` we print out a lot of the other attributes of each object. We do this using the custom function (we wrote it, it wasn't part of the standard library or the `"google.golang.org/cloud/storage"` package) ... we do this using the custom function `statFiles`.

Of course, we are getting **ALL** of the objects in a bucket.

To only get **SOME** of the objects in a bucket, we will need to use a `storage.Query`.

# Using A Query For Specific Results

## MaxResults ( storage.Query )

You will notice that when we used `List` previously, we passed in `nil` as one of the arguments. Here is the definition of the `list` method:

"List lists objects from the bucket. You can specify a query to filter the results. If q is nil, no filtering is applied."

```go
func (b *BucketHandle) List(ctx context.Context, q *Query) (*ObjectList, error)
```

So when we call `List` we can pass in a `*storage.Query`.

We do this in our `query-maxresults` file:

```go
func (d *demo) statFiles() {
	io.WriteString(d.res, "\nRETRIEVING FILE STATS\n")

	client, err := storage.NewClient(d.ctx)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}
	defer client.Close()

	// create a query
	q := storage.Query{
		MaxResults: 2,
	}

	// instead of nil
	// now passing in a *storage.Query
	objs, err := client.Bucket(gcsBucket).List(d.ctx, &q)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		d.dumpStats(obj)
	}
}
```

When we run this code, we will only get two results back.


## Next ( storage.ObjectList )

In our next example `query-maxresults_next` we use the `Next` field from the `*storage.ObjectList` struct (the type which was returned when we ran `List`).

This allows us to *"page"* through our results, seeing two results at a time. You can see this in action in the `query-maxresults_next` example.

The description of `Next` is this:

```go
	// Next is the continuation query to retrieve more
    // results with the same filtering criteria. If there
    // are no more results to retrieve, it is nil.
```

I think it's interesting to look at the relationships of the `List` method, the `*storage.Query` struct, and the `*storage.ObjectList`.

The `List` method takes a `*storage.Query`.

The `List` method returns a `*storage.ObjectList`.

A `*storage.ObjectList` has a field `Next` which is of type `*storage.Query`.

We can then take the value `Next` from our `*storage.ObjectList`, call `List` again, and use the value `Next` as the query argument to provide to `List`.

Pretty cool.

We can learn more about the type `storage.Query` by [looking at it in the docs](https://godoc.org/google.golang.org/cloud/storage#Query):

```go
type Query struct {
    // Delimiter returns results in a directory-like fashion.
    // Results will contain only objects whose names, aside from the
    // prefix, do not contain delimiter. Objects whose names,
    // aside from the prefix, contain delimiter will have their name,
    // truncated after the delimiter, returned in prefixes.
    // Duplicate prefixes are omitted.
    // Optional.
    Delimiter string

    // Prefix is the prefix filter to query objects
    // whose names begin with this prefix.
    // Optional.
    Prefix string

    // Versions indicates whether multiple versions of the same
    // object will be included in the results.
    Versions bool

    // Cursor is a previously-returned page token
    // representing part of the larger set of results to view.
    // Optional.
    Cursor string

    // MaxResults is the maximum number of items plus prefixes
    // to return. As duplicate prefixes are omitted,
    // fewer total results may be returned than requested.
    // The default page limit is used if it is negative or zero.
    MaxResults int
}

```

So in addition to `MaxResults` there are also some other fields which we can use. 

Let's look at `Prefix` next.

## Prefix ( storage.Query)

Adding the `Prefix` field to your `*storage.Query` allows you to select **only the files that have this prefix.**

For example, if we had these files in our bucket on GCS:

```
bar
bar/1
bar/2
boo/
foo/boo/foo1
foo1
foo2
```

And we ran a query with the prefix `foo`, then we would only get these files back from `List`:

```
foo/boo/foo1
foo1
foo2
```

Please notice that we have not used the `Delimiter` field from the query. Because we have not specified a `Delimiter`, our query treats the entire string of the object's name as the object's name. It pays no attention to any delimiters such as the forward slash. That forward slash could just be another letter. The forward slash makes no difference in the object's name when we run a query unless was specify a `Delimiter` in the query, which is what we're going to do next.
 
## Delimiter ( storage.Query )

When you use a delimiter in your query ...

```go
	query := &storage.Query{
		Delimiter: "/",
	}
```

... you get every object whose name does not contain the prefix.

For instance, if we had all of these objects ...

```
2d2542c1d2b96b196df6efbd24e24af88765178e.txt
3d79b3c8f88125cdc1fa1e0b82953460508f79e5.jpeg
48cabad77dc220c4ca91dba582b0b681aa0195fe.md
bar
bar/1
bar/2
bar/nonce/8
bar/nonce/9
bar/nonce/compadre/10
bar/nonce/compadre/11
boo/
boo/yah5
compadre/amigo/diaz6
compadre/luego/hasta7
foo/boo/foo/4
foo/boo/foo3
foo1
foo2
```

... and we ran `List(ctx, query)` with the above query, we would receive these objects ...

```
2d2542c1d2b96b196df6efbd24e24af88765178e.txt
3d79b3c8f88125cdc1fa1e0b82953460508f79e5.jpeg
48cabad77dc220c4ca91dba582b0b681aa0195fe.md
bar
foo1
foo2
```

You can see the above example in the **"query-delimeter"** folder.

What the delimiter is saying is that the delimiter has significance in the name of the object - and thus treat the delimiter as if it has significance.

The significance with which the delimeter is treated is this: it's as if we now have files stored in folders:

```
folder1/folder2/objectname
```

So if we had these files ...

```
firstobject
folder1/folder2/secondobject
```

... and we ran `List(ctx, query)` with this query ...
 
 ```go
 	query := &storage.Query{
 		Delimiter: "/",
 	}
 ```
 
 ... we would receive only this object ...
 
 ```
 firstobject
 ```
 
You can run this example in the **"02"** folder inside the **"query-delimeter"** folder.

To get the secondobject, we would need to run either this query ...
 
```go
 	query := &storage.Query{
 		Prefix: "folder1/folder2/",
 	}
 ```
 
( You can run the above example in the **"01"** folder inside the **"query-prefix-delimeter"** folder. )
 
 ... or this query ...

( You can run the below example in the **"02"** folder inside the **"query-prefix-delimeter"** folder. )
 
 ```go
  	query := &storage.Query{
  		Delimiter: "/",
  		Prefix: "folder1/folder2/",
  	}
  ```
  
If we used the above query with these files in our bucket ...

```
firstobject
folder1/folder2/secondobject
folder1/folder2/folder3/thirdobject
folder1/folder2/folders-can-be-any-string/fourthobject
folder1/folder2/folder3/folder4/fifthobject
```

... we would still get the same results, as you can see in the **"03"** folder inside the **"query-prefix-delimeter"** folder. Now, however, we would be able to access `Prefixes`: the `*storage.ObjectList` which `List` returns would have `Prefixes` field which is a `[]string` that would have values in it. Running the above query on the above files would give us these `Prefixes`:

```
[folder1/folder2/folder3/ folder1/folder2/folders-can-be-any-string/]
```

That is, we would have this `Prefix` ...

```
folder1/folder2/folder3/
```

... and this `Prefix` ...

```
folder1/folder2/folders-can-be-any-string/
```

We could then use those `Prefixes` to access yet more objects in our bucket.

To summarize what we have learned about the `Prefix` and `Delimeter` fields in a `storage.Query` struct:

1. If run a query with a `Prefix`, all objects with that prefix are returned.
1. If run a query with a `Prefix` and a `Delimiter`
  1. All objects with that prefix are returned *as long as they do not have another delimeter after the prefix.*
  1. *If an object does have another delimiter after the prefix*, then we are going to have some additional prefixes returned to us in the `*storage.ObjectList` which `List` returns ( the `Prefixes` field of the `*storage.ObjectList` which is a `[]string` ). The prefixes returned will be everything up to the next delimeter.
1. Our query, which we pass into the `List` method, will return a `*storage.ObjectList` which is a struct with three fields: `Results`, `Next`, and `Prefixes`. We can use `Prefixes` to access additional objects.

Here is how the [documentation](https://godoc.org/google.golang.org/cloud/storage#Query) talks about it with a few slight modifications from me:

```go
type Query struct {
    // Delimiter returns results in a directory-like fashion.
    
    // Results will contain only objects whose names, aside from the
    // prefix, do not contain delimiter. 
    
    // Objects whose names, aside from the prefix, contain another delimiter 
    // will have their name truncated after the next delimiter and returned in prefixes.
    
    // Duplicate prefixes are omitted.
    
    // Optional.
    Delimiter string

    // Prefix is the prefix filter to query objects whose names begin with this prefix.
    // Optional.
    Prefix string

    // Versions indicates whether multiple versions of the same
    // object will be included in the results.
    Versions bool

    // Cursor is a previously-returned page token
    // representing part of the larger set of results to view.
    // Optional.
    Cursor string

    // MaxResults is the maximum number of items plus prefixes to return. 
    // As duplicate prefixes are omitted,
    // fewer total results may be returned than requested.
    // The default page limit is used if it is negative or zero.
    MaxResults int
}
```

To reinforce building queries with a `Prefix` and/or `Delimeter`, two additional examples have been included: the **"04"** and **"05"** folders inside the **"query-prefix-delimeter"** folder.

# ACL - Access Control List

Google Cloud Storage uses access control lists (ACLs) to manage **bucket** and **object** access. ACLs are the mechanism you use to share objects with other users and allow other users to access your buckets and objects.

*An ACL consists of one or more entries, where each entry grants permissions to a scope.* 

**Permissions** define the actions that can be performed against a bucket or object (for example, read or write).
 
 **Scope** defines who the permission applies to (for example, a specific user or group of users). Scopes are sometimes referred to as *grantees.* 

You can read more about ACL's in the **"01_readme"** folder in the **"ACL"** folder of the examples.

In our first code example in the **"02_allusers_rolereader"** folder in the **"ACL"** folder, we see how to set permissions for all users to the role of reader:

```go
	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	wc.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
```

## ACLRule

An `ACLRule` is a [struct](https://godoc.org/google.golang.org/cloud/storage#ACLRule) with two fields: 
 
 ```go
 type ACLRule struct {
     Entity ACLEntity
     Role   ACLRole
 }
 ```
 
### ACLEntity

An `ACLEntity` refers to a user or group. They are sometimes referred to as grantees.
 
It could be in the form of: "user-<userId>", "user-<email>", "group-<groupId>", "group-<email>", "domain-<domain>" and "project-team-<projectId>".
 
Or one of the predefined constants: `AllUsers`, `AllAuthenticatedUsers`.

  ```go
 const (
     AllUsers              ACLEntity = "allUsers"
     AllAuthenticatedUsers ACLEntity = "allAuthenticatedUsers"
 )
 ```
 
When looking at the constants, directly above, some people aren't sure how to read that code. All we're doing there is declaring a constant. To declare a constant in Go, you do it [like this](https://play.golang.org/p/Okk1cj_Wi5):

```go
	const name string = "Henry"
	fmt.Println(name)
```

There is the **const** part, and then the **identifier**, and then the **type**, and then you set it equal to some **value**:

```go
	const identifier type = value
```

You could also do it [like this](https://play.golang.org/p/7RlENlV97X) and leave the type part out:

```go
	const name = "Henry"
	fmt.Println(name)
```

### ACLRole
 
An `ACLRole` is the level of access to grant.

```go
const (
    RoleOwner  ACLRole = "OWNER"
    RoleReader ACLRole = "READER"
)
```

You can see what ACL's look like, by default, when you don't specifically set any ACL's by running the code in the **"03_default"** folder in the **"ACL"** folder of the examples.


# Delete

You may have already seen this method in our code:

```go
func (d *demo) delFiles() {

	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		if err := d.bucket.Object(obj.Name).Delete(d.ctx); err != nil {
			log.Errorf(d.ctx, "deleteFiles: unable to delete bucket %q, file %q: %v", d.bucket, obj.Name, err)
			return
		}
	}
}
```

To delete an object, you specify the bucket and the object, and then you call the Delete method.

In the [documentation](https://godoc.org/google.golang.org/cloud/storage#ACLRole), you would go these functions and methods:

```go
func NewClient(ctx context.Context, opts ...cloud.ClientOption) (*Client, error)

func (c *Client) Bucket(name string) *BucketHandle

func (b *BucketHandle) Object(name string) *ObjectHandle

func (o *ObjectHandle) Delete(ctx context.Context) error
```

`NewClient` gives you a pointer to a client. You can now call the method `Bucket` which gives you a pointer to a `BucketHandle`. You can now call the method `Object` which gives you a pointer to an `ObjectHandle`. You can now call the method `Delete`.

# Two Examples

There are two last examples: **file-browser** and **gcs-example** for you to peruse.

The file-browser example was written by Caleb Doxsey. 

The gcs-example was written by Google.

# A Challenge

Store images in GCS for three different users where the object name on GCS is in the format <username>/<object-name>.txt. For each user, print the file names.

You can find a solution to this challenge using images instead of text files in the **challenge** folder.