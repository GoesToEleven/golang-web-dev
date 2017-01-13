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

I'm putting Go at root

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
nano ~/.profile
```

.bashrc is for Bash only
.profile is used by other shells as well

## Add this to .profile at the top of the file

```
export GOROOT=/go
export GOPATH=$HOME/goworkspace
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
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
go get github.com/GoesToEleven/go-hello-world 
```

# go install your project

Navigate to the directory with the go code:

```
go install
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

```
docker run -d --name hello -p 80:8080 user/image-name ~/goworkspace/bin/go-hello-world
```

-p maps network ports; docker host port : container port