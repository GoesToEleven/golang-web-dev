# Running containers

Running an application inside a container takes a single command: 

```
docker run
```

## Run ubuntu

```
docker run ubuntu /bin/echo 'Hello world'
```

```docker run``` runs a container.

```ubuntu``` is the image you run, for example the Ubuntu operating system image. When you specify an image, Docker looks first for the image on your Docker host. If the image does not exist locally, then the image is pulled from the public image registry Docker Hub.

```/bin/echo``` is the command to run inside the new container.

The container launches. Docker creates a new Ubuntu environment and executes the ```/bin/echo``` command inside it and then prints out:

Hello world

## Run ubuntu tty

```
docker run -t -i ubuntu /bin/bash
```

In this example:

1. ```docker run``` runs a container.
1. ```ubuntu``` is the image you would like to run.
1. -t flag assigns a pseudo-tty or terminal inside the new container.
1. -i flag allows you to make an interactive connection by grabbing the standard input (STDIN) of the container.
1. ```/bin/bash``` launches a Bash shell inside our container.

The container launches. We can see there is a command prompt inside it. Try running some Linux commands: pwd, ls, ....

Type ```exit``` to exit the tty session.

## Run an ubuntu daemon

```
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
```

-d flag runs the container in the background (to daemonize it).
-c is the command to run 

[-c string	If the -c flag is present then commands are read from string.](http://heirloom.sourceforge.net/sh/sh.1.html)

 
 ### See running containers (processes)
 
```
docker ps
```
 
 ## View container's standard output
 
 Now, we know the container is running. But is it doing what we asked it to do? 
 
 To see this we’re going to look inside the container using the docker logs command.
 
 ```docker logs``` shows the standard output of a container.

```
docker logs <container name>
```

## Stop container

```
docker stop <container name>
```

## Verify container is stopped

```
docker ps
docker ps -a
```
