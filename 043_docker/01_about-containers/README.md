# Applications & machines

Back in the day, an application would be created and run on a machine. There was a ratio of 1:1. Each application ran on its own machine.
 
When wondering what type of a machine to purchase for their application, developers always purchased the most powerful machine possible. There was no certainty as to how much an application might be in demand. To cover their asses, they errored on the side of too much power instead of not enough.

The result of this: many machines were greatly underutilized. Many machines only ran at a small fraction of their capacity.

# Virtual machines

There were many machines with unused capacity.

Virtual machines found a way to use that capacity.

One physical machine could host several virtual machines.

Each virtual machine appeared just like a real machine to its user.

Here is an image from [docker's website](https://www.docker.com/what-docker) that shows the architecture of virtual machines:

![Virtual Machines](vm.png)

There was a downside to this, though: each virtual machine had its own operating system. This created license and maintenance costs, as well as using up resources on the physical machine.

The question arose: Is there a way to divide up a physical machine, but only have one operating system on that machine?

# Containers

Linux is built in such a way that this is possible: you can create separate user spaces. Each user space has its own file system and processes. The user spaces are all segregated and separate from each other.

Containers leverage this technology of the Linux operating system.

A container is just like a virtual machine WITHOUT THE OPERATING SYSTEM.

Each PHYSICAL MACHINE has ONE LINUX OPERATING SYSTEM.

Containers use the Linux OS of the physical machine. 

To the user, each container appears just like a real machine to its user.

Here is an image from [docker's website](https://www.docker.com/what-docker) that shows a comparison between virtual machines and containers:

![Virtual Machines vs Containers](vm-containers.png)

# Docker

Docker is the most popular container system.

You build an **IMAGE** and then create as many **CONTAINERS** as you would like from that image.

You can then run those containers on physical machines.

Docker containers run on Linux, Mac OS, and Windows.

