

# Setup systemd
1. sudo nano /etc/systemd/system/goserver.service
```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/ubuntu/webapp/hello-world
WorkingDirectory=/home/ubuntu/webapp/public
User=ubuntu
Group=ubuntu
Restart=always

[Install]
WantedBy=multi-user.target
```
1. sudo systemctl enable goserver.service
1. sudo systemctl start goserver.service
1. sudo systemctl status goserver.service