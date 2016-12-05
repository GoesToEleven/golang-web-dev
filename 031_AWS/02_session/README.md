# Deploying our session example

1. Create your binary
  - GOOS=linux GOARCH=amd64 go build -o [some-name] *.go

1. Copy you "templates" folder to the server
  - scp -i /path/to/[your].pem templates/* ubuntu@[public-IP]:/home/ubuntu/templates

1. If necessary, SSH into your server
  - ssh -i /path/to/[your].pem ubuntu@[public-IP]:

1. Run your code
  - sudo ./[some-name]
  - check it in a browser at [public-IP]:8080