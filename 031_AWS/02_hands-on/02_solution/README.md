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

# Persisting your application

  To run our application after the terminal session has ended, we must do the following:

  1. Create a configuration file
    - sudo nano /etc/systemd/system/```<filename>```.service

  ```
  [Unit]
  Description=Go Server

  [Service]
  ExecStart=/home/<username>/<path-to-exe>/<exe>
  WorkingDirectory=/home/<username>/<exefolderpath>
  User=<username>
  Group=<username>
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