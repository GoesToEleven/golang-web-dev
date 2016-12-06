# Setting up

[Setting Up for Amazon RDS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html)

1. [Sign Up for AWS](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.SignUp)
1. [Create an IAM User](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.IAM)
1. [Determine Requirements](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.Requirements)
1. [Provide Access to the DB Instance in the VPC by Creating a Security Group1.](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html#CHAP_SettingUp.SecurityGroup)

# Create an IAM User

Services in AWS, such as Amazon RDS, require that you provide credentials when you access them, so that the service can determine whether you have permission to access its resources. 

We recommend that you use AWS Identity and Access Management (IAM). 

Create an IAM user, and then add the user to an IAM group with administrative permissions. 

If you signed up for AWS but have not created an IAM user for yourself, you can create one using the IAM console.

## Create an IAM user for yourself and add the user to an Administrators group
   
 1. Sign in to the Identity and Access Management (IAM) console at https://console.aws.amazon.com/iam/.
   
 1. In the navigation pane, choose Users, and then choose Add user.
   
 1. For User name, type a user name, such as Administrator. The name can consist of letters, digits, and the following characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-). The name is not case sensitive and can be a maximum of 64 characters in length.
   
 1. Select the check box next to AWS Management Console access, select Custom password, and then type the new user's password in the text box. You can optionally select Require password reset to force the user to select a new password the next time the user signs in.
   
 1. Choose Next: Permissions.
   
 1. On the Set permissions for user page, choose Add user to group.
   
 1. Choose Create group.
   
 1. In the Create group dialog box, type the name for the new group. The name can consist of letters, digits, and the following characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-). The name is not case sensitive and can be a maximum of 128 characters in length.
   
 1. For Filter, choose Job function.
   
 1. In the policy list, select the check box for AdministratorAccess. Then choose Create group.
   
 1. Back in the list of groups, select the check box for your new group. Choose Refresh if necessary to see the group in the list.