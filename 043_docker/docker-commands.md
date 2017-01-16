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

## xx05_push-pull

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

# xx06_daemon

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

# xx07_custom-image

docker run -t -i ubuntu /bin/bash
apt-get update
apt-get -y upgrade
apt-get install curl
curl --head www.google.com
docker ps
docker commit <ID of running container> <create a new name for your new image>
docker run -t -i <name of your new container> /bin/bash
curl --head www.google.com

# xx08_add-golang
docker run -t -i ubuntu /bin/bash
apt-get update
apt-get -y upgrade
apt-get install curl
apt-get install nano
apt-get install git
curl -O <current version of Go for Linux download link>
tar -xvf <file name of the tar.gz file that was downloaded with curl>
mv go /
cd
mkdir goworkspace
cd goworkspace
mkdir bin pkg src
cd
nano ~/.profile

## Add this to .profile at the top of the file
```
export GOROOT=/go
export GOPATH=$HOME/goworkspace
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

source ~/.profile 
go version
Mac: command + t
docker ps
docker commit <ID of running container> <create a new name for your new image>
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker login
docker push <docker hub username>/<image name>

# xx09_add-binary
go get -d github.com/GoesToEleven/go-hello-world 
go install
cd ~/goworkspace/bin
chmod 777 go-hello-world
docker ps
docker commit <ID of running container> <create a new name for your new image>
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
docker login
docker push <docker hub username>/<image name>
docker run -d -p 80:8080 toddmcleod/helloworld  /root/goworkspace/bin/helloworld






















