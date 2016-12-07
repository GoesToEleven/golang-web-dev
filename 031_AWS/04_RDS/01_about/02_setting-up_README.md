# Setting up

[Setting Up for Amazon RDS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html)
1. [Sign Up for AWS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.SignUp)1. [Create an IAM User](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.IAM)1. [Determine Requirements](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.Requirements)1. [Provide Access to the DB Instance in the VPC by Creating a Security Group1.](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.SecurityGroup)

# Create an IAM User

Services in AWS, such as Amazon RDS, require that you provide credentials when you access them, so that the service can determine whether you have permission to access its resources. 

We recommend that you use AWS Identity and Access Management (IAM). 

Create an IAM user, and then add the user to an IAM group with administrative permissions. 

If you signed up for AWS but have not created an IAM user for yourself, you can create one using the IAM console.

## Create an IAM user for yourself and add the user to an Administrators group
   
 1. [Sign in to the Identity and Access Management (IAM)](console at https://console.aws.amazon.com/iam/)
   
 1. In the navigation pane, choose **Users**, and then choose **Add user**.
   
 1. For **User name**, type a user name, such as **Administrator**. The name can consist of letters, digits, and the following characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-). The name is not case sensitive and can be a maximum of 64 characters in length.
   
 1. Select the check box next to **AWS Management Console access**, select **Custom password**, and then type the new user's password in the text box. You can optionally select Require password reset to force the user to select a new password the next time the user signs in.
   
 1. Choose **Next: Permissions**.
   
 1. On the **Set permissions for user** page, choose **Add user to group**.
   
 1. Choose **Create group**.
   
 1. In the **Create group** dialog box, type the name for the new group such as **dbAdmins**. The name can consist of letters, digits, and the following characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-). The name is not case sensitive and can be a maximum of 128 characters in length.
   
 1. For **Filter**, choose **Job function**.
   
 1. In the **policy type list**, select the check box for **DatabaseAdministrator**. Then choose Create group.
   
 1. Back in the list of groups, select the check box for your new group. Choose **Refresh** if necessary to see the group in the list.
 
 1. Choose Next: **Review** to see the list of group memberships to be added to the new user. When you are ready to proceed, choose **Create User**.
 
 1. **Download user credentials**. This will be the last time the credentials can be downloaded. Store them in a safe place which you will remember.
 
 You can use this same process to create more groups and users, and to give your users access to your AWS account resources. To learn about using policies to restrict users' permissions to specific AWS resources, go to Access Management and Example Policies for Administering AWS Resources.
 
 To sign in as this new IAM user, sign out of the AWS console, then use the following URL, where your_aws_account_id is your AWS account number without the hyphens (for example, if your AWS account number is 1234-5678-9012, your AWS account ID is 123456789012):
 
 https://your_aws_account_id.signin.aws.amazon.com/console/
 Enter the IAM user name and password that you just created. When you're signed in, the navigation bar displays "your_user_name @ your_aws_account_id".
 
 If you don't want the URL for your sign-in page to contain your AWS account ID, you can create an account alias. From the IAM dashboard, click Customize and enter an alias, such as your company name. To sign in after you create an account alias, use the following URL:
 
 https://your_account_alias.signin.aws.amazon.com/console/
 To verify the sign-in link for IAM users for your account, open the IAM console and check under AWS Account Alias on the dashboard.
 
 ## Determine Requirements
 
 The basic building block of Amazon RDS is the DB instance. 
 
 The DB instance is where you create your databases. 
 
 A DB instance provides a network address called the **Endpoint**. 
 
 Your applications connect to the **endpoint** exposed by the DB instance whenever they need to access the databases created in that DB instance. 
 
 The information you specify when you create the DB instance controls configuration elements such as storage, memory, database engine and version, network configuration, security, and maintenance periods.
 
 You must know your DB instance and network needs before you create a security group and before you create a DB instance. For example, you must know the following:
 
1. What are the memory and processor requirements for your application or service?
1. Your DB instance is most likely in a **virtual private cloud (VPC)**; some legacy instances are not in a VPC. The security group rules you need to connect to a DB instance depend on whether your DB instance is in a **default VPC**, in a **user-defined VPC**, or **outside of a VPC**. The follow list describes the rules for each VPC option:
  - **Default VPC** — If you specify the default VPC when you create the DB instance:
    - You must create a VPC security group that authorizes connections from the application or service to the Amazon RDS DB instance with the database. Note that you must use the Amazon EC2 API or the Security Group option on the VPC Console to create VPC security groups. For information, see Step 4: Create a VPC Security Group.
1. You must specify the default DB subnet group. If this is the first DB instance you have created in the region, Amazon RDS will create the default DB subnet group when it creates the DB instance.
1. User-defined VPC — If you want to specify a user-defined VPC when you create a DB instance:
1. You must create a VPC security group that authorizes connections from the application or service to the Amazon RDS DB instance with the database. Note that you must use the Amazon EC2 API or the Security Group option on the VPC Console to create VPC security groups. For information, see Step 4: Create a VPC Security Group..
 The VPC must meet certain requirements in order to host DB instances, such as having at least two subnets, each in a separate availability zone. For information, see Amazon RDS and Amazon Virtual Private Cloud (VPC).
 You must specify a DB subnet group that defines which subnets in that VPC can be used by the DB instance. For information, see the DB Subnet Group section in Working with a DB Instance in a VPC.
 No VPC — if your AWS account does not have a default VPC, and you do not specify a user-defined VPC:
 You must create a DB security group that authorizes connections from the devices and Amazon RDS instances running the applications or utilities that will access the databases in the DB instance. For more information, see Working with DB Security Groups.
 Do you need failover support? On Amazon RDS, a standby replica of your DB instance that can be used in the event of a failover is called a Multi-AZ deployment. If you have production workloads, you should use a Multi-AZ deployment. For test purposes, you can usually get by with a single instance, non-Multi-AZ deployment.
 Does your AWS account have policies that grant the permissions needed to perform Amazon RDS operations? If you are connecting to AWS using IAM credentials, your IAM account must have IAM policies that grant the permissions required to perform Amazon RDS operations. For more information, see Authentication and Access Control for Amazon RDS.
 What TCP/IP port will your database be listening on? The firewall at some companies may block connections to the default port for your database engine. If your company firewall blocks the default port, choose another port for the new DB instance. Note that once you create a DB instance that listens on a port you specify, you can change the port by modifying the DB instance.
 What region do you want your database in? Having the database close in proximity to the application or web service could reduce network latency.
 What are your storage requirements? Do you need to use Provisioned IOPS? Amazon RDS provides three storage types: magnetic, General Purpose (SSD), and Provisioned IOPS (input/output operations per second) . Magnetic storage, also called standard storage, offers cost-effective storage that is ideal for applications with light or burst I/O requirements. General purpose, SSD-backed storage, also called gp2, can provide faster access than disk-based storage. Provisioned IOPS storage is designed to meet the needs of I/O-intensive workloads, particularly database workloads, that are sensitive to storage performance and consistency in random access I/O throughput. For more information on Amazon RDS storage, see Storage for Amazon RDS.
 Once you have the information you need to create the security group and the DB instance, continue to the next step.