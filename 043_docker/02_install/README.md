# Overview

This is what we will do:

1. install Docker
1. run a software image in a container
1. locate an interesting image on Docker Hub
1. run the image on your own machine
1. modify an image to create your own and run it
1. create a Docker Hub account and repository
1. push your image to Docker Hub for others to share

# Install

[Install docker](https://docs.docker.com/)

Choose **Docker engine / install**

# Verify installation

Check the docker version
```
docker version
```

Run a sample "hello world" application
```
docker run hello-world 
```

# Images, Repository, Containers

You use an **image** to make a **container**.

You can make an unlimited number of **containers** from one **image**.

An **image** is stored in a docker image **repository**. 

**Docker hub** is one docker image repository you can use.

See all containers (processes) on your machine
```
docker ps -a
```

See all running containers on your machine
```
docker ps
```