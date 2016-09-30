# Understanding templates

A template allows us to create one document and then merge data with it.

We are learning about templates so that we can create one document, a web page, and then merge customized data to that page.

Web templates allow us to serve personalized results to users.

Think of Facebook - you log into your main page and see results tailored for you. That main page was created once. It is a template. However, for each user, that template gets populated with data specific to that user.

Another common exposure to templates that most of us get every day - junk mail.

A company creates a piece of mail to send to everyone, and then they merge data with that template to customize the mailing for each individual. The result:

## Template Example - Merged With Data

*** 

Dear Mr. Jones,
 
Are you tired of high electric bills?

We have noticed that your house at .....

*** 

Dear Mr. Smith,
 
Are you tired of high electric bills?

We have noticed that your house at .....

***

## Template Example - The Template

Dear {{Name}},

Are you tired of high electric bills?

We have noticed that your house at .....
