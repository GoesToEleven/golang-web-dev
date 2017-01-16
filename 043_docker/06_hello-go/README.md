# Create Go code

Create a directory

Create main.go

Create code for a "hello world" web app

# Create .dockerignore file

If you want to ignore any files or directories in your build, add a .dockerignore file

# Create Dockerfile

The Dockerfile must be named Dockerfile.

The Dockerfile will include EVERYTHING in the current directory, and descendent directories, in the image which is built (unless told to ignore something by the .dockerignore file)

The Dockerfile may start with a comment

# Yo, this is my Dockerfile, Yo

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

```


Create Dockerfile

## Create a directory

```
mkdir mydockerbuild2
```

## Enter directory

```
cd mydockerbuild2
```

## Create a docker build file

```
nano Dockerfile
```

## Add these lines to your build file

```
# Creates an ubuntu image with curl installed
FROM ubuntu:latest
RUN apt-get -y update && apt-get install -y curl
```

# Build an image from the Dockerfile

## Docker build
Build the image using the docker build command. 

The -t parameter gives your image a tag, so you can run it more easily later. 

Don’t forget the . command, which tells the docker build command to look in the current directory for a file called Dockerfile. 

```
docker build -t curler .
```

The above command reads the Dockerfile in the current directory and processes its instructions one by one to build an image called **curler** on your local machine.

# Run your new docker-whale

See what images you have on your machine:

```
docker images
```

Run your image

```
docker run -it curler /bin/bash
```

run a curl command

```
curl --head www.google.com
```
