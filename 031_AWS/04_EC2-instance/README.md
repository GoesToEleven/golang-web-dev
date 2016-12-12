## Amazon EC2 instance

1. create a key pair
1. create an instance
1. test your instance

# Create a key pair

We will create a “key pair” to allow a user to SSH into our virtual machines.

Public key cryptography is any cryptographic system that uses pairs of keys: **public keys** which may be disseminated widely, and **private keys** which are known only to the owner. 

This accomplishes two functions: 

  - authentication: the public key is used to verify that a holder of the paired private key sent the message
  - encryption: only the holder of the paired private key can decrypt the message encrypted with the public key

1. EC2
1. key pair
1. create
  - give it a name that is descriptive, for example: **devs-spring19**

# Create an instance
1. EC2 / instances / launch instance
1. choose an AMI (Amazon Machine Instance)
  - amazon linux AMI (this includes PostgreSQL)
1. choose instance size
1. configure instance
  - stick with defaults
1. storage
1. tag
  - value: web-server-0001
1. configure security group
  - choose the "web-servers-sg" security group we created
1. launch
  - specify "key pair" we want the instance to use
1. launch instance

# Test your instance

## SSH
1. either way:
  // TODO  updated ssh to this
  - ssh -i "devs-spring19.pem" ec2-user@ec2-52-53-172-37.us-west-1.compute.amazonaws.com
  // from this
  - ssh -i ~/Downloads/devs-spring19.pem ec2-usr@```<web server's public IP address>```
  - example: ssh -i ~/Downloads/devs-spring19.pem ec2-usr@52.53.172.37

## HTTP
1. helloworld app
  - access denined until: add rule to the "web-servers-sg" security group
    - HTTP TCP 80 My IP
1. endpoints app
  - ping endpoint: have the route "/ping" return "OK"
  - instance endpoint: have the route "/instance" return the EC2 instance