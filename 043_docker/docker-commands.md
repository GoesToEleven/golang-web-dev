# These are the docker commands we have run

docker version

docker run hello-world 

docker ps -a

docker ps

docker run docker/whalesay cowsay boo

docker images

FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay

docker build -t docker-whale .

docker images

docker run docker-whale

docker images

docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>

docker images

docker login

docker push <docker hub username>/<image name>

docker images

docker rmi - f <image ID or image name>

docker --help
docker rmi --help

docker run <yourusername>/docker-whale






