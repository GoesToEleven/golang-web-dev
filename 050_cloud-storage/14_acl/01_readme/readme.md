# Documentation About ACL

[About Access Control Lists](https://cloud.google.com/storage/docs/access-control#About-Access-Control-Lists)

# Access Control List

Google Cloud Storage uses access control lists (ACLs) to manage bucket and object access. ACLs are the mechanism you use to share objects with other users and allow other users to access your buckets and objects.

**An ACL consists of one or more entries, where each entry grants permissions to a scope.** 

**Permissions** define the actions that can be performed against a bucket or object (for example, read or write).
 
 **Scope** defines who the permission applies to (for example, a specific user or group of users). Scopes are sometimes referred to as *grantees.* 
 
 The maximum number of ACL entries you can create for a bucket or object is 100. When the ACL scope is a group or domain, it counts as one ACL entry regardless of how many users are in the group or domain.

When a user requests access to a bucket or object, the Google Cloud Storage system reads the bucket or object ACL and determines whether to allow or reject the access request. 

If the ACL grants the user permission for the requested operation, the request is allowed. If the ACL does not grant the user permission for the requested operation, the request fails and a **403 StatusForbidden** error (Access Denied) is returned.

# Permissions - Object

**Reader**

Lets a user download an object's data.

**Writer**

You *cannot apply this* permission to objects.

**Owner**

Gives a user READER access. It also lets a user read and write object metadata, including ACLs.

**Default**

Objects have the predefined [project-private ACL](https://cloud.google.com/storage/docs/access-control#predefined-project-private) applied when they are uploaded. Objects are always owned by the original requester who uploaded the object.

# Permissions - Bucket

**Reader**

Lets a user list a bucket's contents.

The following bucket metadata properties are not returned with a bucket's resource without OWNER: acl, cors, defaultObjectAcl, lifecycle, logging, owner, and projectNumber. 

**Writer**

Lets a user list, create, overwrite, and delete objects in a bucket.

The following bucket metadata properties cannot be changed: acl, cors, defaultObjectAcl, lifecycle, logging, versioning and website.

**Owner**

Gives a user READER and WRITER permissions on the bucket. It also lets a user read and write bucket metadata, including ACLs.

**Default**

Buckets have the predefined [project-private ACL](https://cloud.google.com/storage/docs/access-control#predefined-project-private) applied when they are created. Buckets are always owned by the project-owners group.

# [Scopes](https://cloud.google.com/storage/docs/access-control#scopes)

**Google Storage ID**

A Google Storage ID is a string of 64 hexadecimal digits that identifies **a specific Google account** holder or a specific [Google group](https://groups.google.com/forum/#!overview). It is sometimes referred to as a canonical ID. The following is an example of a Google Storage ID:

```
84fac329bceSAMPLE777d5d22b8SAMPLE77d85ac2SAMPLE2dfcf7c4adf34da46
```

Project teams are identified by a Google Storage ID. The project editors group and project owners group are also identified using Google Cloud Storage IDs. These IDs are unique to a project.

**Google account email address**

Every user who has **a Google account** must have a unique email address associated with that account. You can specify a scope by using any email address that is associated with a Google account, such as a gmail.com address.

Google Cloud Storage remembers email addresses as they are provided in ACLs until the entries are removed or overwritten. If a user changes email addresses, you should update ACL entries to reflect these changes.

**Google group email address**

Every Google group has a unique email address that is associated with the group. For example, the Google Cloud Storage Announce has the following email address: gs-announce@googlegroups.com. You can find the email address that is associated with a Google group by clicking About on the homepage of every Google group. For more information about Google groups, see the Google groups homepage.

Like Google account email addresses, Google Cloud Storage remembers group email addresses as they are provided in ACLs until the entries are removed or overwritten. You do not need to worry about updating Google Group email addresses because Google Group email addresses are permanent and unlikely to change.


**Google Apps domain**

[Google Apps](https://apps.google.com/) customers can associate their email accounts with an Internet domain name. When you do this, each email account takes the form username@yourdomain.com. You can specify a scope by using any Internet domain name that is associated with a Google Apps account.


**Special identifier for all Google account holders**

This special scope identifier represents anyone who is authenticated with a Google account. The special scope identifier for all Google account holders is `allAuthenticatedUsers`.

**Special identifier for all users**

This special scope identifier represents anyone who is on the Internet, with or without a Google account. The special scope identifier for all users is `allUsers`.

****