# Overview

We are going to add a database to our application.

To do this well, we need to configure AWS in a certain way.

## Here are the steps

1. create EC2 security groups
1. create an ELB load balancer
1. add EC2 instance to load balancer
1. create an AMI (Amazon Machine Image)
1. create auto scaling


# Create security groups

A security group acts as a virtual *firewall* that controls the traffic for one or more instances.

When you launch an instance, you associate one or more security groups with the instance.

You add rules to each security group that allow traffic to or from its associated instances.

You can modify the rules for a security group at any time; the new rules are automatically applied to all instances that are associated with the security group. When we decide whether to allow traffic to reach an instance, we evaluate all the rules from all the security groups that are associated with the instance.

## ELB (load balancer) security group
1. add this rule
  - HTTP TCP 80 Anywhere
1. copy *Group ID*
  
## Web tier security group
1. add these rules
  - HTTP TCP 80 Custom IP ```<load-balancer-sg Group ID>```
  - SSH TCP 22 My IP
1. copy **Group ID**
1. add this rule
  - MySql TCP 3306 Custom IP ```<web-servers-sg Group ID>```

# Load balancer
1. EC2 / Load balancers / Create load balancer
  - application or classic
  - name: web-elb
  - http & https
  - default VPC
  - add two subnets
1. configure security groups
  - choose "load-balancer" security group which we just setup
1. configure routing
  - target group: web-servers-tg1
  - ping path: /ping
  - allows us to define a "healthy" web server
  - load balancer will only forward to healthy web servers
1. register targets
  - we don't have any yet; we'll add them after we create them
1. create

# Add EC2 instance to load balancer
1. EC2/ load balancers / instances / edit instances
1. add the web server we have created / save
  - it takes a moment before it's "in service"; try hitting refresh
1. description / DNS name / enter this into a browser to see your load balancer in action

# Create an AMI (Amazon Machine Image)
1. EC2 / Instances / right-click instance / create image
  - image name: web-architecture-2019-10-31
  - description: web server 2019 October 31
  - no reboot: unchecked 
    - allowing your instance to reboot gives a better image
1. create image

## Launch a new instance of your AMI in a new availability zone (AZ)
1. what AZ is your current instance running in?
  - EC2 / instances / look at the availability zone and make note of it
1. launch a new instance from your AMI
  - EC2 / AMIs / right click / launch / next: configure
1. subnet: ```<choose a different AZ>``` / next: storage / next
1. tag
  - value: web-server-0002
1. security group
  - choose the "web-servers-sg" security group we created
1. launch
  - specify "key pair" we want the instance to use
1. launch instance

## test
  - copy the public ip then go to these routes in your browser
    - ```<public ip>```/ping
    - ```<public ip>```/instance

## Add new EC2 instance to load balancer
1. EC2/ load balancers / instances / edit instances
1. add the web server we have created / save
  - it takes a moment before it's "in service"; try hitting refresh
1. description / DNS name / enter this into a browser to see your load balancer in action
  - refresh your browser to see the switching between web-servers-sg
  
# Create auto scaling

Auto Scaling helps you maintain application availability and allows you to scale your Amazon EC2 capacity up or down automatically according to conditions you define. You can use Auto Scaling to help ensure that you are running your desired number of Amazon EC2 instances. Auto Scaling can also automatically increase the number of Amazon EC2 instances during demand spikes to maintain performance and decrease capacity during lulls to reduce costs. Auto Scaling is well suited both to applications that have stable demand patterns or that experience hourly, daily, or weekly variability in usage. 

## Configure auto scaling

1. EC2 / autoscaling / launch configuration
1. create auto scaling group / create launch configuration
1. My AMIs / choose your AMI
  - my image name was "web-architecture-2019-10-31"
  - next / next
1. configure details
  - name: **auto-scale-config-2019-10-31**
  - next / next
1. configure security group
  - select an existing security group / select the "web-servers-sg" security group
  - next / next / create launch configuration
  - choose an existing key pair / create launch configuration
  
## Create auto scaling group

1. Configure auto scaling group
  - name: **auto-scale-group-2019-10-31**
  - group size: this is the minimum number of instances we'll always be running
  - network: default vpc
  - subnet: choose the availability zones (AZs) into which you've launched instances
  - advanced details
    - load balancing: check "receive traffic from elastic load balancer"
    - select your load balancer
    - health check: ELB (this is what we set up)
1. configure scaling policies
  - keep group at initial size
1. configure tags
    - value: web-server-auto-scaled
1. create auto scaling group
1. Scaling policies
  - this is where we'd add policies to say when we scale up / scale down 
