# Preview

We will search docker hub for an image. We will then run that image.

When you run an image, that image is copied to your local machine.

When you run an image in a container, Docker downloads the image to your computer. This local copy of the image saves you time. Docker only downloads the image again if the image’s source changes on the hub. You can, of course, delete the image yourself. You’ll learn more about that later. 

# Locate whalesay

Locate the **docker/whalesay** image on [Docker Hub](https://hub.docker.com/)

You can also do this to search docker hub:

```
docker search <search term>
```

Caveat: This is code. When you grab someone else's code to use with your code, make sure you know what's in it.

# Run the program

```
docker run docker/whalesay cowsay boo
```

# See the images on your machine

```
docker images
```

# See all containers (processes) on your machine

```
docker ps -a
```

```
docker ps
```

# Reflection

You searched for an image on Docker Hub. 

You used your command line to run an image. Effectively you ran a piece of Linux software on whatever machine you're using (Windows, Mac, Linux). 

You learned that running an image copies it on your computer.