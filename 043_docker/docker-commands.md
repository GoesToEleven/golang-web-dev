# These are the docker commands we used

## 02_install

docker version
docker run hello-world 
docker ps -a
docker ps

## 03_run-images

docker search <search term>
docker run docker/whalesay cowsay boo
docker images
docker ps -a
docker ps

## 04_build-image

mkdir mydockerbuild
cd mydockerbuild
nano Dockerfile

FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay

docker build -t docker-whale .
docker images
docker run docker-whale

# 05_curl

mkdir mydockerbuild2
cd mydockerbuild2
nano Dockerfile

```
# Creates an ubuntu image with curl installed
FROM ubuntu:latest
RUN apt-get -y update && apt-get install -y curl
```

docker build -t curler .
docker images
docker run -it curler /bin/bash
curl --head www.google.com

# 06_hello-go

```
# Some comment
FROM golang:1.8-onbuild
MAINTAINER youremail@gmail.com
```

docker build -t my-app .
docker run -d -p 80:80 my-app
 
# 07_push-pull

docker images
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker login
docker push <docker hub username>/<image name>
docker images
docker rmi -f <image ID or image name>
docker --help
docker <COMMAND> --help
docker rmi --help
docker run <yourusername>/<app-name>























