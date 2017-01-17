# Build an image

1. Create a Dockerfile
1. Build an image from the Dockerfile

# Create Dockerfile

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

Don’t forget the ```.``` command, which tells the docker build command to look in the current directory for a file called Dockerfile.

```
docker build -t <give it an image name> .
```

The above command reads the Dockerfile in the current directory and processes its instructions one by one to build an image called <whatever name you gave it> on your local machine.

# Run your new image

See what images you have on your machine:

```
docker images
```

## Run your image

### Mac / Linux / maybe Windows
```
MAC / LINUX: docker run -it <image name> /bin/bash
```

### Windows: if you get this error ```the input device is not a TTY. If you are using mintty, try prefixing the command with "winpty" ``` use the below code

```
WINDOWS: winpty docker run -it <image name> bash
```

## run a curl command

```
curl --head www.google.com
```

# Exit

Exit your image by typing ```exit```
