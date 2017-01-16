# Build an image

1. Create a Dockerfile
1. Build an image from the Dockerfile

# Create Dockerfile

## Create a directory

```
mkdir mydockerbuild
```

## Enter directory

```
cd mydockerbuild
```

## Create a docker build file

```
nano Dockerfile
```

## Add these lines to your build file

```
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay
```

### FROM
The FROM keyword tells Docker which image your image is based on. 

### RUN
The RUN statement will install the fortunes program into the image. 

The whalesay image is based on Ubuntu, which uses apt-get to install packages. 

These two commands refresh the list of packages available to the image and install the fortunes program into it. 

The fortunes program prints out wise sayings for our whale to say.

### CMD
The CMD statement tells the image the final command to run after its environment is set up. 

This command runs fortune -a and pipes its output to the cowsay command.


# Build an image from the Dockerfile

## Docker build
Build the image using the docker build command. 

The -t parameter gives your image a tag, so you can run it more easily later. 

Don’t forget the . command, which tells the docker build command to look in the current directory for a file called Dockerfile. 

```
docker build -t docker-whale .
```

The above command reads the Dockerfile in the current directory and processes its instructions one by one to build an image called **docker-whale** on your local machine.

# Run your new docker-whale

See what images you have on your machine:

```
docker images
```

Run your image

```
docker run docker-whale
```

Run it again, and again, and again to see different output.
