# Create a new EC2 instance

Use the Amazon Linux AMI 

[EC2 Docker Instructions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html)

change the permissions on your pem if you created a new one
```
sudo chmod 400 your.pem
```

# Connect to your machine

ssh into your instance
```
ssh -i /path/to/[your].pem ec2-user@[public-DNS]
```

# Update the installed packages and package cache on your instance.
```
sudo yum update -y
```

# Install Docker.
```
sudo yum install -y docker
```

# Start the Docker service.
```
sudo service docker start
```

# Add the ec2-user to the docker group so you can execute Docker commands without using sudo.
```
sudo usermod -a -G docker ec2-user
```

# Log out and log back in again to pick up the new docker group permissions.

# Verify that the ec2-user can run Docker commands without sudo.
```
docker info
```

In some cases, you may need to reboot your instance to provide permissions for the ec2-user to access the Docker daemon. Try rebooting your instance if you see the following error:
``` 
Cannot connect to the Docker daemon. Is the docker daemon running on this host?
```

# Run the Go app from the previous example
```
docker run -d -p 80:80 toddmcleod/golang-hello-world
```

# Check that it's running
```
docker ps
```

# View the web app from a browser
```
Use the IP address of your instance
```