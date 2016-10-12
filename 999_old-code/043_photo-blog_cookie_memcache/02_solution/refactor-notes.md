# Set Memcache

### Cookie Returned To Client

Before we write a cookie back to the client, we always call `func makeCookie`

This is a good point to then also call a func to put our data into memcache.

This way, the value stored in the cookie will be the same as the value stored in memcache.

```go
func makeCookie(mm []byte, id string, req *http.Request) *http.Cookie {

	// Anytime a cookie is created, let's print the id
	// The id is the key for the value in memcache
	// Having the id will allow us to lookup the value in memcache
	log.Fatalln("ID:", id)

	b64 := base64.URLEncoding.EncodeToString(mm)

	// SEND DATA TO BE STORED IN MEMCACHE
	storeMemc([]byte(b64), id, req)

	// SEND DATA TO BE STORED IN COOKIE
	code := getCode(b64) // hmac
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}
```

```go
func storeMemc(bs []byte, id string, req *http.Request) {
	ctx := appengine.NewContext(req)
	item1 := memcache.Item{
		Key:   id,
		Value: bs,
	}
	memcache.Set(ctx, &item1)
}
```

Wherever `func makeCookie` is called, we will need to update our code to ensure a value of type `*http.Request` is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

Good to note: memcache is key:value storage. Our key is the uuid for the user. Our value is the same value we are writing to our cookie. Remember, our cookie has `id|value|hmac code`. In both the cookie and memcache, the value is the `model` data marshaled to `JSON` and then encoded to `base64`.

Remember this: 

**model --> JSON --> base64**

To unwind this, we will need to do this:

**base64 --> JSON --> model**

# FYI, This Is An Unrealistic Example

FYI, this is an unrealistic example though it is building block in your educational process. 

Our uuid which uniquely identifies a user is stored in the cookie. Our []string which stores the paths to a user's photos are also stored in that cookie. 

We are now also storing all of that data in memcache. 

We access that data in memcache by the user's uuid (memcache stores key:value pairs). 

Well, to have the uuid, we have to have the cookie. 

And if we have the cookie, we have the []string with paths to the user's photos. 

**So why also get the []string from memcache? We already have it!**

Well, we're getting it from memcache just to learn this process. 

Eventually we will store our data in the datastore. We will have our uuid in the cookie. We will then check memcache for the []string which stores photo paths. If it's not in memcache, we will then check the datastore for this []string.

Eventually, we will also store the user's photos in google cloud storage (our hard drive in the cloud). 

So the whole process, at the end of this will be:

1. store uuid in **cookie**
1. store user session info in **memcache**
1. store user session info and user info in **datastore**
1. store user files in **google cloud storage**
1. attempt to retrieve user session info from **memcache**
  1. if unable to retrieve user session info from **memcache**, retrieve user session info from **datastore**
    1. store this session info in **memcache**
    1. next we retrieve user session info, it's in **memcache**
1. retrieve user photos from **google cloud storage**

### IMPORTANT! - This Code Runs But Is Not Functional

We will be running this code on Google App Engine.

We are not yet storing photos in Google Could Storage.

As our code just tries to write the photos to our server, that write will fail.

This is still, however, a necessary stepping stone in our learning process.

# Retrieve Data From Memcache

### Change func Model

The function `func Model` returns the a value of type model - **this has all of the user's session data**.

We wil need to change `func Model` to see if we have data in memcache and, if so, to then get the data from there.

To do this, we will need to have access to the `*http.Request` as this is necessary to access memcache.

Let's modify `func Model` to have a parameter of type `*http.Request` ...

```go
// Model returns a value of type model
func Model(c *http.Cookie, req *http.Request) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Fatalln("Error decoding base64", err)
	}

	m := unmarshalModel(bs)

	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != nil {
		m.Pictures = m2.Pictures
		log.Fatalln("PICTURE PATHS RETURNED FROM MEMCACHE")
		log.Fatalln(m.Pictures)
	}

	return m
}
```

Wherever func Model is called, we will need to update our code to ensure a value of type `*http.Request` is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

### If There Is Data In Memcache ...

Now, anytime `func Model` is called, it will check to see if there is data in memcache and, if so, it will use the `[]string` picture paths from that data.
 
 ```go
func retrieveMemc(req *http.Request, id string) model {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, id)

	// decode item.Value from base64
	bs, err := base64.URLEncoding.DecodeString(string(item.Value))
	if err != nil {
		log.Fatalln("Error decoding base64 in retrieveMemc", err)
	}

	// unmarshal from JSON
	var m model
	if item != nil {
		m = unmarshalModel(bs)
	}
	return m
}
 ```

### Refactored / Abstracted Code

Modularized code in `func Model` and put it in `func unmarshalModel`  

```go
func unmarshalModel(bs []byte) model {

	var m model
	err := json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	return m
}
```
# Upload Photos

When we upload photos, we will need to also add the new file information to our paths of photos which we are storing in the model as type `[]string`.

To increase code readability, there is some redundant code in here which could be refactored.

```go
func addPhoto(fName string, c *http.Cookie, req *http.Request) *http.Cookie {
	xs := strings.Split(c.Value, "|")

	// memcache
	id := xs[0]
	m2 := retrieveMemc(req, id)
	m2.Pictures = append(m2.Pictures, fName)
	mm := marshalModel(m2)
	b64 := base64.URLEncoding.EncodeToString(mm)
	storeMemc([]byte(b64), id, req)

	// cookie
	usrData := xs[1]
	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Fatalln("Error decoding base64", err)
	}
	m := unmarshalModel(bs)
	m.Pictures = append(m.Pictures, fName)
	mm = marshalModel(m)
	b64 = base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64) // hmac
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}
```

# Refactor Code For Appengine

`package main` ... to ... `package mem`

I could have called `package mem` something else like, oh, I don't know, maybe `package mickeymouse`

Took code out of `func main` and put it into `func init`

Added `app.yaml` file

# Note the use of appengine/log

```go
google.golang.org/appengine/log
```

```go
ctx := appengine.NewContext(req)
log.Errorf(ctx, "Error decoding base64: %s", err)
```

Read more about it here:

https://godoc.org/google.golang.org/appengine/log

# Run Your App

Enter this at the terminal: `goapp serve`

Go to this url: `http://localhost:8080`

Look in your terminal and get the id

Go to this url: `http://localhost:8000/memcache`

Enter the id and see what's in memcache.