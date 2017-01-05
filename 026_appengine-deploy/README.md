# buying a domain

https://domains.google/#/

# deploying to [Google Cloud](https://cloud.google.com/)
- [install google appengine](https://cloud.google.com/appengine/docs/go/download)
- [make sure python is installed VERSION 2.7.x](https://www.python.org/downloads/release/python-2712/)
  - python -V
  - configure environment PATH variables
- google cloud developer console
  - create a project
  - get the project ID
- update the app.yaml file with your project ID

```
application: temp-137512
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```
- deploy to that project
```
appcfg.py -A <YOUR_PROJECT_ID> -V v1 update .
```
- view your project
  - http://YOUR_PROJECT_ID.appspot.com/


example
http://temp-145415.appspot.com/


change DNS info at google domains
point your domain to your appengine project
