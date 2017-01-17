# Create Go app

Create a directory

Create main.go

Create code for a "hello world" web app

# Create .dockerignore file

If you want to ignore any files or directories in your build, add a .dockerignore file

# Create Dockerfile

The Dockerfile must be named Dockerfile.

The Dockerfile will include EVERYTHING in the current directory, and descendent directories, in the image which is built (unless told to ignore something by the .dockerignore file)

The Dockerfile may start with a comment

```
# Yo, this is my Dockerfile, Yo
```

The Dockerfile must have FROM as the first instruction

The FROM says what image your are building this image from.

In most cases, you will start with an image to build your image.

You then later new images on top of your starting FROM image, and that all finally becomes your finished image.

Images are made out of images layered on top of images.

You can add in a MAINTAINER instruction if you'd like and say who built this image

We are going to build our image FROM a golang image so ...

Go to docker hub

Search for golang

Find the golang image you want

We will use: golang:1.8-onbuild

So our Dockerfile will be

```
# Some comment
FROM golang:1.8-onbuild
MAINTAINER youremail@gmail.com
```

# Now build our image

```
docker build -t my-app .
```

-t means "tag" or "give it a name" 

The name we gave it is "my-app"

The dot "." means the code for this image is in this current directory

Make sure you are in the correct directory when you run this
 
# Now create a container from your image and run it
 
```
 docker run -d -p 80:80 my-app
```
 
-d means run this detached, as a daemon, eg, not dependent on the terminal session

-p means map ports; mapping ```<host machine port>:<to docker container port>```
 
 
# Verify it's running

Go to your browser and see if it's running