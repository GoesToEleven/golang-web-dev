Take the code from "030_sessions/08_expire-session" and get it running on AWS.

Remember to change the port number from 8080 to 80.

You will need to modify your service configuration file to include information about your working directory:

```
[Unit]
Description=Go Server

[Service]
ExecStart=/home/<username>/<exepath>
WorkingDirectory=/home/<username>/<working-dir>
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```