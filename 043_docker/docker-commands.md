# These are the docker commands we used

## 02_install

docker version
docker run hello-world 
docker ps -a
docker ps

## 03_run-images

docker search
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

## 05_push-pull

docker images
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker images
docker login
docker push <docker hub username>/<image name>
docker images
docker rmi -f <image ID or image name>
docker --help
docker rmi --help
docker run <yourusername>/docker-whale

# 06_daemon

docker run
docker run ubuntu /bin/echo 'Hello world'
docker run -t -i ubuntu /bin/bash
exit
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
docker ps
docker logs <container name>
docker stop <container name>
docker ps
docker ps -a






