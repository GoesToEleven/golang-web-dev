# Run an ubuntu container

```
docker run -t -i ubuntu /bin/bash
```

# Update apt-get

```
apt-get update
apt-get -y upgrade
```

# Install curl

```
apt-get install curl
```

# Test curl

```
curl --head www.google.com
```

# Create image from container

You **cannot** exit your current container.

If you exit your current container, all changes to that container will be lost.

Launch a new tab in your terminal window. Alternatively, you can detach from your current container with these commands:  ```ctrl+p, ctrl+q```

At the terminal, find out the running docker containers:

```
docker ps
```

Now create a new image from that running container:

```
docker commit <ID of running container> <create a new name for your new image>
```

# Verify your new container

```
docker run -t -i <name of your new container> /bin/bash
```

```
curl --head www.google.com
```