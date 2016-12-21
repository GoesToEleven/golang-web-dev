# Create auto scaling

Auto Scaling helps you maintain application availability and allows you to scale your Amazon EC2 capacity up or down automatically according to conditions you define. You can use Auto Scaling to help ensure that you are running your desired number of Amazon EC2 instances. Auto Scaling can also automatically increase the number of Amazon EC2 instances during demand spikes to maintain performance and decrease capacity during lulls to reduce costs. Auto Scaling is well suited both to applications that have stable demand patterns or that experience hourly, daily, or weekly variability in usage.

## Configure auto scaling

1. EC2 / autoscaling / launch configuration
1. create auto scaling group / create launch configuration
1. My AMIs / choose your AMI
  - next / next
1. configure details
  - name: *auto-scale-config-2019-10-31*
  - next / next
1. configure security group
  - select an existing security group / select the "web-tier" security group
  - next / next / create launch configuration
  - choose an existing key pair / create launch configuration

## Create auto scaling group

1. Configure auto scaling group
  - name: *auto-scale-group-2019-10-31*
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













