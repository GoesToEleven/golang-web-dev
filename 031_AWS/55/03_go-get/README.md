# Install Go

choose the "source" code not an installer or archive file

terminal: 
1. wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
  - Go downloads
1. tar -xzf go1.7.4.linux-amd64.tar.gz
  - Go is unpacked
1. rm -rf go1.7.4.linux-amd64.tar.gz
  - removes the tar file
1. pwd
  - print current working directoy
1. mkdir goworkspace
1. cd gowoworkspace
1. mkdir bin pkg src
1. cd ../


# Add environment variables
1. nano .bashrc
  - add this to your .bashrc file:

export GOROOT=/home/ubuntu/go
export GOPATH=/home/ubuntu/goworkspace
export PATH=$PATH:/home/ubuntu/goworkspace/bin
export PATH=$PATH:/home/ubuntu/go/bin

# Confirm setup
1. source ~/.bashrc
  - refreshes bash environment variables
1. go version
  - confirms you have go installed

# go get code

# Go run main.go