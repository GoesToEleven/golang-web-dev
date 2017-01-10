# Run an ubuntu container

```
docker run -t -i ubuntu /bin/bash
```

# Update apt-get

```
apt-get update
apt-get -y upgrade
```

# Get Curl

```
apt-get install curl
```

# Install nano

```
apt-get install nano
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
go env
```

# Build a go application

Build a go application.

See "08_hello-docker" for the binary I created




# go get that go application

