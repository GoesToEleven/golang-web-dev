# Overview

We are going to add a database to our application.

To do this well, we need to configure AWS in a certain way.

## Here are the steps

1. create an IAM user
** 1. create EC2 security groups
**** 1. create an ELB load balancer
***** 1. add EC2 instance to load balancer
* 1. create a database
*** 1. create an AMI (Amazon Machine Image)
****** 1. create auto scaling

# Create an IAM user

The account we have been using is our AWS master account.

A better security practice is to 
  - set up a new **user**
  - give that user **permissions and priveleges**
  
We don't want to use the master account because that account has access to everything, including billing information. It is good to get into the practice of setting up users with permissions.

The master account should only be used to manage our payment information for AWS.

We should create an account for everyone else, including ourselves.

1. IAM (identity access management)
  - users / add user / enter user name(s) for example: **ob1**
  - AWS Management Console access / custom password for example: **theforce**
  - uncheck: require password reset
1. permissions
  - attach existing policies directly / AdministratorAccess

## Logging in: IAM User Sign-in URL

A user needs to use a specific URL to sign-in. You can find this URL on the dashboard for IAM.
  - you can create an alias for the account which makes it easier to remember and use
  
1. Log out from the master account
1. Go to your specific user sign-in URL
    - example in video: https://starwars007.signin.aws.amazon.com/console
1. enter the username and password for the user
1. Try going to billing information for the account
  - you won't be able to access it because you're not the master account user



# Create security groups

A security group acts as a virtual **firewall** that controls the traffic for one or more instances.

When you launch an instance, you associate one or more security groups with the instance.

You add rules to each security group that allow traffic to or from its associated instances.

You can modify the rules for a security group at any time; the new rules are automatically applied to all instances that are associated with the security group. When we decide whether to allow traffic to reach an instance, we evaluate all the rules from all the security groups that are associated with the instance.

## ELB (load balancer) security group
  - allow HTTP port 80 accessible from anywhere
1. EC2 / security groups / create security group
  - name: load-balancer-sg
  - description: allows access from anywhere on HTTP port 80
  - VPC: default
1. **Add rule**
  - HTTP TCP 80 Anywhere
  - create
1. copy **Group ID** from the "load-balancer-sg" security group we just created
  - we will need it to define that the web servers should only accept traffic from the load-balancer
    - group id:
  
## Web tier security group
  - allow HTTP port 80 accessible from ELB (load balancer)
  - allow MySQL/TCP port 3306 only from web servers
1. EC2 / security groups / create security group
  - name: web-servers-sg
  - description: allows access from ELB load-balancer
  - VPC: default
1. **Add rule**
  - HTTP TCP 80 Custom IP ```<load-balancer-sg Group ID>```
  - create
1. copy **Group ID** from the "web-servers-sg" security group we just created
  - we will need it to define another rule
1. Inbound
1. **Add rule**
  - MySql TCP 3306 Custom IP ```<web-servers-sg Group ID>```
1. **Add rule**
  - SSH TCP 22 My IP

Any service belonging to the load-balancer-sg security group can accept traffic from:
  - the load-balancer **OR** 
  - TCP port 3306 traffic from services belonging to "web-servers-sg" security group **OR**
  - an EC2 instance (web server) belonging to the "web-servers-sg" security group **OR**
  - SSH from a specific IP

# Load balancer
1. EC2 / Load balancers / Create load balancer
  - application or classic
  - name: web-elb
  - http & https
  - default VPC
  - add two subnets
1. configure security groups
  - choose "load-balancer-sg" security group which we just setup
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

# Create a database
1. RDS / intances / launch DB instance
1. choose database engine: Aurora
  - production or development/learning
1. details
  - defaults except for
    - storage: 5GM
    - db instance identifier: db-test
    - master username: **admin**
    - master password: **password**
1. advanced settings
  - defaults except for
    - security group: "web-server"
    - database name: database-test
1. Launch DB instance
1. test that you can connect to the database
  - connect to the **endpoint** of your database
    ```
    mysql -h <your database endpoint up to the .com but not the port> -u admin -p
    
    show databases;
    
    quit
    ```

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
