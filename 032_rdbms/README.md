1. install mysql
2. install workbench
3. create mysql db on aws
  - important fix to avoid connection problem (10060 Error): "add a rule for the inbound TCP port in the security group used by your database instance to communicate with your client workstation. From the DB Instance details page in the RDS console, select the instance, click on the detail tab then click on the Security Group link for the DB Instance. You should be taken to the Security Group section on the EC2 dashboard. Select the Security Group in the list and at the bottom of the page, select the Inbound tab. Click Edit and a popup will appear. Click Add Rule and select MySQL/Aurora for the Type. For the source, you should be able to select "My IP" and the IP address of your client workstation should be populated. After you apply this, you should be able to connect." ([source](https://forums.aws.amazon.com/thread.jspa?threadID=231166))
4. connect workbench to rds mysql db
5. https://www.youtube.com/watch?v=k68Y-XYapEI
