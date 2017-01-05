# Cross-site scripting (XSS)

Cross-site scripting (XSS) is a type of computer security vulnerability typically found in web applications.

XSS enables attackers to inject client-side scripts into web pages viewed by other users. 

A cross-site scripting vulnerability may be used by attackers to bypass access controls such as the [same-origin policy](https://en.wikipedia.org/wiki/Same-origin_policy): you have a script on one site that makes a request to another site. For example: you come to my cool website about kittens, and a script runs to transfer money from UnionBank to my foreign account. If it wasn't for the "same-origin policy" implemented in browsers, and if you had a cookie on your machine that said you were logged into Union Bank, then the money would transfer. 

Cross-site scripting carried out on websites accounted for roughly 84% of all security vulnerabilities documented by Symantec as of 2007. Their effect may range from a petty nuisance to a significant security risk, depending on the sensitivity of the data handled by the vulnerable site and the nature of any security mitigation implemented by the site's owner.

***

# Same-origin policy

In computing, the same-origin policy is an important concept in the web application security model. 

Under the policy, a web browser permits scripts contained in a first web page to access data in a second web page, but only if both web pages have the same origin. 

An origin is defined as a combination of URI scheme, hostname, and port number. 

This policy prevents a malicious script on one site from obtaining access to sensitive data on another site.

## Example

Assume a user is visiting a banking website and doesn't log out. 

Then he goes to another site and that site has some malicious JavaScript code running in the background that requests data from the banking site. 

Because the user is still logged in on the banking site, without the "same-origin policy" implemented in browsers, that malicious code could do anything on the banking site. 

For example, get a list of your last transactions, create a new transaction, etc. This is because the browser can send and receive session cookies to the banking website based on the domain of the banking website. A user visiting that malicious site would expect that the site he is visiting has no access to the banking session cookie. While this is true, the JavaScript has no direct access to the banking session cookie, but it could still send and receive requests to the banking site with the banking site's session cookie, essentially acting as a normal user of the banking site! 

Regarding the sending of new transactions, even CSRF (cross-site request forgery) protections by the banking site have no effect, because the script can simply do the same as the user would do. 

So this is a concern for all sites where you use sessions and/or need to be logged in. 

## All modern browsers implement some form of the Same-Origin Policy as it is an important security cornerstone.

This mechanism bears a particular significance for modern web applications that extensively depend on HTTP cookies to maintain authenticated user sessions, as servers act based on the HTTP cookie information to reveal sensitive information or take state-changing actions. 

A strict separation between content provided by unrelated sites must be maintained on the client-side to prevent the loss of data confidentiality or integrity.