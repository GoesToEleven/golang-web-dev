# Create an instance

1. Create an AWS Account
  - enter credit card info

1. Sign into console

1. Create a new EC2 instance
  - services / EC2
  - launch instance
  - choose your instance
  - add storage / 30GB free
  - add tags / webserver
  - security / ssh / http
  - launch
  - create new key pair / download

# Deploy your binary

1.
  - mv [src] [dst] / sudo chmod 400 your.pem

1. Build hello world
  - GOOS=linux GOARCH=amd64 go build

1. Copy your binary to the sever
  - scp -i /path/to/[your].pem ./main ec2-user@[public-DNS]:
    - "ec2-user" might be "ubuntu" depending upon your machine
  - say "yes" to The authenticity of host ... can't be established.

1. SSH into your server
  - ssh -i /path/to/[your].pem ec2-user@[public-DNS]

1. Run your code
  - sudo chmod 700 mybinary
  - sudo ./mybinary
  - check it in a browser at [public-IP]

1. Exit
  - ctrl + c
  - exit

# Persisting your application

To run our application after the terminal session has ended, we must do one of the following:

## Possible options
1. screen
1. init.d
1. upstart
1. system.d

## System.d
1. Create a configuration file
  - cd /etc/systemd/system/
  - sudo nano ```<filename>```.service

```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/<username>/<exepath>
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

1. Add the service to systemd.
  - sudo systemctl enable ```<filename>```.service
1. Activate the service.
  - sudo systemctl start ```<filename>```.service
1. Check if systemd started it.
  - sudo systemctl status ```<filename>```.service
1. Stop systemd if so desired.
  - sudo systemctl stop ```<filename>```.service


# Troubleshooting

A possible issue could be that you're cross-compiling for the wrong architecture: AWS might have assigned you a different machine than the one used in this example. To solve this problem, we will install Go on the AWS machine and then run "go env" to see GOOS & GOARCH for that machine.

1. download Go
  - wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
1. unpack go
  - tar -xzf go1.7.4.linux-amd64.tar.gz
1. remove the tar file
  - rm -rf go1.7.4.linux-amd64.tar.gz
1. make your go workspace
  - mkdir goworkspace
  - cd gowoworkspace
  - mkdir bin pkg src
  - cd ../
1. add environment variables
  - nano .bashrc
```
export GOROOT=/home/ubuntu/go
export GOPATH=/home/ubuntu/goworkspace
export PATH=$PATH:/home/ubuntu/goworkspace/bin
export PATH=$PATH:/home/ubuntu/go/bin
```
1. refresh environment variables
  - source ~/.bashrc
1. confirm installation
  - go version
1. get machine GOOS & GOARCH info
  - go env

# Troubleshooting

Sometimes students miss setting port openings in security. If you are having issues, check to make sure these settings are correct - and please note, you IP address for SSH will either be 0.0.0.0/0 or something different than mine.
![](security.png)

