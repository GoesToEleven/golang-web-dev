# Deploying our session example

1. change your port number from 8080 to 80

1. create your binary
  - GOOS=linux GOARCH=amd64 go build -o [some-name]

1. SSH into your server
  - ssh -i /path/to/[your].pem ubuntu@[public-DNS]:

1. create directories to hold your code
  - for example, "wildwest" & "wildwest/templates"

1. copy binary to the server

1. copy you "templates" to the server
  - scp -i /path/to/[your].pem templates/* ubuntu@[public-DNS]:/home/ubuntu/templates

1. chmod permissions on your binary

1. Run your code
  - sudo ./[some-name]
  - check it in a browser at [public-IP]

# Persisting your application

  To run our application after the terminal session has ended, we must do the following:

  1. Create a configuration file
    - sudo nano /etc/systemd/system/```<filename>```.service

  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/<username>/<path-to-exe>/<exe>
  WorkingDirectory=/home/<username>/<exe-working-dir>
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








# FOR EXAMPLE
  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/ubuntu/cowboy
  WorkingDirectory=/home/ubuntu
  User=root
  Group=root
  Restart=always

  [Install]
  WantedBy=multi-user.target
```








