# Run an ubuntu container

```
docker run -t -i ubuntu /bin/bash
```

# Update apt-get

```
apt-get update
apt-get -y upgrade
```

# Install curl

```
apt-get install curl
```

# Install nano

```
apt-get install nano
```

# Install git

```
apt-get install git
```

# Get Go

```
curl -O <current version of Go for Linux download link>
```

# Unpack Go

```
tar -xvf <file name of the tar.gz file that was downloaded with curl>
```

# Move go

I'm putting Go at "/". 

Move it there if it's not already there.

```
mv go /
```

# Create Go workspace

1. create a dir called "goworkspace" (or whatever else you want to call it)
1. inside goworkspace, create these dirs: bin, pkg, src

```
cd
mkdir goworkspace
cd goworkspace
mkdir bin pkg src
```

# Set environment variables 

```
cd
nano ~/.profile
```

.bashrc is for Bash only
.profile is used by other shells as well

## Add this to .profile at the top of the file

```
export GOROOT=/go
export GOPATH=$HOME/goworkspace
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

## Reload environment variables

```
source ~/.profile 
```

## Verify installation

```
go version
```

# Commit your container to an image

You **cannot** exit your current container.

If you exit your current container, all changes to that container will be lost.

Launch a new tab in your terminal window. Alternatively, you can detach from your current container with these commands:  ```ctrl+p, ctrl+q```

At the terminal, find out the running docker containers:

Mac: command + t

```
docker ps
```

Now create a new image from that running container:

```
docker commit <ID of running container> <create a new name for your new image>
```

# Go get your project

Switch back to your original container.

```
go get -d github.com/GoesToEleven/go-hello-world 
```

-d flag = just download

# go install your project

Navigate to the directory with the go code:

```
go install
```

# Change permissions

```
cd ~/goworkspace/bin
chmod 777 go-hello-world
```

# Commit your container to an image

Launch a new tab in your terminal window. Alternatively, you can detach from your current container with these commands:  ```ctrl+p, ctrl+q```

At the terminal, find out the running docker containers:

```
docker ps
```

```
docker commit <ID of running container> <create a new name for your new image>
```

# Push your image

```
docker tag <image ID>  <docker hub username>/<image name>:<version label or tag>
```

```
docker login
```

```
docker push <docker hub username>/<image name>
```

# Run a container from your image

docker run -d -p <host port>:<container port> <username>/<image name>  <commands to run in the container>

```
docker run -d -p 80:8080 toddmcleod/helloworld  /root/goworkspace/bin/helloworld
```

The -d flag runs the container in the background (as a so-called daemon).

The -p flag maps any required network ports inside the container to your host. This lets you view the web application. <host port>:<container port>

The toddmcleod/helloworld image is a pre-built image that contains a simple Go web application.

The remaining arguments make up the command that is run inside the container.
















