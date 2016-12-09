# About RDS

[Amazon Relational Database Service (Amazon RDS)](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html)

Amazon Relational Database Service (Amazon RDS) is a web service that makes it easier to set up, operate, and scale a relational database in the cloud. It provides cost-efficient, resizeable capacity for an industry-standard relational database and manages common database administration tasks.

# Components

[Amazon RDS Components](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html#Welcome.Concepts)

## DB Instances

The basic building block of Amazon RDS is the DB instance.
 
A DB instance is an isolated database environment in the cloud.  
 
A DB instance **can contain multiple user-created databases,** and you can access it by using the same tools and applications that you use with a stand-alone database instance. 
 
You can create and modify a DB instance by using the **AWS Management Console**.

Each DB instance runs a **DB engine**. 

Amazon RDS currently supports the **MySQL, MariaDB, PostgreSQL, Oracle, and Microsoft SQL Server DB engines**. 

The **computation and memory capacity of a DB instance** is determined by its DB instance class. You can select the DB instance that best meets your needs. If your needs change over time, you can change DB instances. 

You can run a DB instance on a virtual private cloud using Amazon's **Virtual Private Cloud (VPC) service**. When you use a virtual private cloud, you have control over your virtual networking environment: you can select your own IP address range, create subnets, and configure routing and access control lists. 

The basic functionality of Amazon RDS is the same whether it is running in a VPC or not; Amazon RDS manages backups, software patching, automatic failure detection, and recovery. There is no additional cost to run your DB instance in a VPC. 

## Regions and Availability Zones

Amazon cloud computing resources are housed in highly available data center facilities in different areas of the world (for example, North America, Europe, or Asia). **Each data center location is called a region**.

Each region contains multiple distinct locations called **Availability Zones (AZs)**. 

Each Availability Zone is engineered to be isolated from failures in other Availability Zones, and to provide inexpensive, low-latency network connectivity to other Availability Zones in the same region. By launching instances in separate Availability Zones, you can protect your applications from the failure of a single location. 

You can run your DB instance in several Availability Zones, an option called a Multi-AZ deployment. When you select this option, Amazon automatically provisions and maintains a synchronous standby replica of your DB instance in a different Availability Zone. The primary DB instance is synchronously replicated across Availability Zones to the standby replica to provide data redundancy, failover support, eliminate I/O freezes, and minimize latency spikes during system backups.

## Security Groups

A security group **controls the access to a DB instance.** 

Amazon RDS uses **DB security groups**, **VPC security groups**, and **EC2 security groups**. 

In simple terms, a **DB security group** controls access to a DB instance that is not in a VPC, a **VPC security group** controls access to a DB instance inside a VPC, and an Amazon **EC2 security group** controls access to an EC2 instance and can be used with a DB instance. 

## DB Parameter Groups

You **manage the configuration of a DB** engine by using a DB parameter group. A DB parameter group contains engine configuration values that can be applied to one or more DB instances of the same instance type. 

Amazon RDS applies a default DB parameter group if you don’t specify a DB parameter group when you create a DB instance. The default group contains defaults for the specific database engine and instance class of the DB instance.