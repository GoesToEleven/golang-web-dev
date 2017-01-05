# Understanding AWS

Amazon Web Services (AWS)offers a suite of cloud-computing services that make up an on-demand computing platform. These services operate from 14 geographical regions across the world. They include Amazon Elastic Compute Cloud ("EC2") and Amazon Simple Storage Service ("S3"). 

As of 2016 AWS has more than 70 services including compute, storage, networking, database, analytics, application services, deployment, management, mobile, developer tools and tools for the Internet of things. 

Amazon markets AWS as a service to provide large computing capacity quicker and cheaper than building your own physical server farm.

# Services

## EC2 - Elastic Compute Cloud
Amazon Elastic Compute Cloud (EC2) forms a central part of Amazon.com's cloud-computing platform, Amazon Web Services (AWS), by allowing users to **rent virtual computers** on which to run their own computer applications. EC2 encourages scalable deployment of applications by providing a web service through which a user can boot an *Amazon Machine Image (AMI)* to configure a virtual machine, which Amazon calls an *instance*, containing any software desired. A user can create, launch, and terminate server-instances as needed, paying by the hour for active servers – hence the term "elastic". EC2 provides users with control over the geographical location of instances that allows for latency optimization and high levels of redundancy.

### EC2 Key Pairs
Amazon EC2 uses public–key cryptography to encrypt and decrypt login information. Public–key cryptography uses a public key to encrypt a piece of data, such as a password, then the recipient uses the private key to decrypt the data. The public and private keys are known as a key pair. To log in to your instance, you must create a key pair, specify the name of the key pair when you launch the instance, and provide the private key when you connect to the instance. Linux instances have no password, and you use a key pair to log in using SSH. With Windows instances, you use a key pair to obtain the administrator password and then log in using RDP.

### EC2 Security Groups
A security group acts as a *virtual firewall* that controls the traffic for one or more instances. When you launch an instance, you associate one or more security groups with the instance. You add rules to each security group that allow traffic to or from its associated instances. You can modify the rules for a security group at any time; the new rules are automatically applied to all instances that are associated with the security group. When we decide whether to allow traffic to reach an instance, we evaluate all the rules from all the security groups that are associated with the instance.

### EC2 User data
You can run scripts when your instance launches.

### EC2 AMI's
An Amazon Machine Image (AMI) provides the information required to launch an instance, which is a virtual server in the cloud. You specify an AMI when you launch an instance, and you can launch as many instances from the AMI as you need. You can also launch instances from as many different AMIs as you need.

### EC2 Elastic Load Balancer (ELB)
Elastic Load Balancing automatically distributes incoming application traffic across multiple Amazon EC2 instances. It enables you to achieve fault tolerance in your applications, seamlessly providing the required amount of load balancing capacity needed to route application traffic. Elastic Load Balancing offers two types of load balancers that both feature high availability, automatic scaling, and robust security. These include the Classic Load Balancer that routes traffic based on either application or network level information, and the Application Load Balancer that routes traffic based on advanced application level information that includes the content of the request. The Classic Load Balancer is ideal for simple load balancing of traffic across multiple EC2 instances, while the Application Load Balancer is ideal for applications needing advanced routing capabilities, microservices, and container-based architectures. Application Load Balancer offers ability to route traffic to multiple services or load balance across multiple ports on the same EC2 instance.

### EC2 Auto Scaling
Auto Scaling helps you maintain application availability and allows you to scale your Amazon EC2 capacity up or down automatically according to conditions you define. You can use Auto Scaling to help ensure that you are running your desired number of Amazon EC2 instances. Auto Scaling can also automatically increase the number of Amazon EC2 instances during demand spikes to maintain performance and decrease capacity during lulls to reduce costs. Auto Scaling is well suited both to applications that have stable demand patterns or that experience hourly, daily, or weekly variability in usage. 

## Databases

1. RDS
Amazon Relational Database Service (Amazon RDS) makes it easy to set up, operate, and scale a relational database in the cloud. It provides cost-efficient and resizable capacity while managing time-consuming database administration tasks, freeing you up to focus on your applications and business. Amazon RDS provides you six familiar database engines to choose from, including Amazon Aurora, PostgreSQL, MySQL, MariaDB, Oracle, and Microsoft SQL Server.

1. DynamoDB
Amazon DynamoDB is a fast and flexible NoSQL database service for all applications that need consistent, single-digit millisecond latency at any scale. It is a fully managed cloud database and supports both document and key-value store models. Its flexible data model and reliable performance make it a great fit for mobile, web, gaming, ad tech, IoT, and many other applications. According to Forrester research, Amazon's DynamoDB is the most popular NoSQL cloud database.

1. ElastiCache
Amazon ElastiCache is a web service that makes it easy to deploy, operate, and scale an in-memory data store or cache in the cloud. The service improves the performance of web applications by allowing you to retrieve information from fast, managed, in-memory data stores, instead of relying entirely on slower disk-based databases. Amazon ElastiCache supports two open-source in-memory engines:

  - Redis - a fast, open source, in-memory data store and cache. Amazon ElastiCache for Redis is a Redis-compatible in-memory service that delivers the ease-of-use and power of Redis along with the availability, reliability and performance suitable for the most demanding applications. Both single-node and up to 15-shard clusters are available, enabling scalability to up to 3.55 TiB of in-memory data. ElastiCache for Redis is fully managed, scalable, and secure - making it an ideal candidate to power high-performance use cases such as Web, Mobile Apps, Gaming, Ad-Tech, and IoT.

  - Memcached - a widely adopted memory object caching system. ElastiCache is protocol compliant with Memcached, so popular tools that you use today with existing Memcached environments will work seamlessly with the service.

Using Amazon ElastiCache, you can add an in-memory layer to your infrastructure in a matter of minutes.

## S3 - Simple Storage Service
Amazon S3 provides storage through web services interfaces (REST, SOAP, and BitTorrent). Amazon launched S3 in the United States in March 2006 and in Europe in November 2007.

At its inception, Amazon charged end users US$0.15 per gigabyte/month, with additional charges for bandwidth used sending and receiving data, and a per-request (get or put) charge. On November 1, 2008, pricing moved to tiers where end users storing more than 50 terabytes receive discounted pricing. Amazon says that S3 uses the same scalable storage infrastructure that Amazon.com uses to run its own global e-commerce network.

S3 uses include web hosting, image hosting, and storage for backup systems. 

## Route 53
You can use Amazon Route 53 to register new domains, transfer existing domains, route traffic for your domains to your AWS and external resources, and monitor the health of your resources.

## Cloudfront
Amazon CloudFront is a global *content delivery network (CDN)* service that accelerates delivery of your websites, APIs, video content or other web assets. It integrates with other Amazon Web Services products to give developers and businesses an easy way to accelerate content to end users with no minimum usage commitments.

## Cloudwatch
Amazon CloudWatch is a *monitoring service* for AWS cloud resources and the applications you run on AWS. You can use Amazon CloudWatch to collect and track metrics, collect and monitor log files, set alarms, and automatically react to changes in your AWS resources. Amazon CloudWatch can monitor AWS resources such as Amazon EC2 instances, Amazon DynamoDB tables, and Amazon RDS DB instances, as well as custom metrics generated by your applications and services, and any log files your applications generate. You can use Amazon CloudWatch to gain system-wide visibility into resource utilization, application performance, and operational health. You can use these insights to react and keep your application running smoothly.

## Identity Access Management (IAM)
AWS Identity and Access Management (IAM) enables you to securely control access to AWS services and resources for your users. Using IAM, you can create and manage AWS users and groups, and use permissions to allow and deny their access to AWS resources. IAM is a feature of your AWS account offered at no additional charge. You will be charged only for use of other AWS services by your users. To get started using IAM, or if you have already registered with AWS, go to the AWS Management Console and get started with IAM Best Practices.

## Elastic Beanstalk
AWS Elastic Beanstalk is an orchestration service offered from Amazon Web Services for deploying infrastructure which orchestrates various AWS services, including 
- EC2
- S3
- Simple Notification Service (SNS)
- CloudWatch
- autoscaling
- Elastic Load Balancers.

Elastic Beanstalk provides an additional layer of abstraction over the bare server and OS; users instead see a pre-built combination of OS and platform, such as "64bit Amazon Linux 2014.03 v1.1.0 running Ruby 2.0 (Puma)" or "64bit Debian jessie v2.0.7 running Python 3.4 (Preconfigured - Docker)".

Deployment requires a number of components to be defined: an 'application' as a logical container for the project, a 'version' which is a deployable build of the application executable, a 'configuration template' that contains configuration information for both the Beanstalk environment and for the product. 

Finally an 'environment' combines a 'version' with a 'configuration' and deploys them.

Executables themselves are uploaded as archive files to S3 beforehand and the 'version' is just a pointer to this.

Supported software include: 
- PHP
- Java
- .NET
- Node.JS
- Python
- Ruby
- Docker
- Go