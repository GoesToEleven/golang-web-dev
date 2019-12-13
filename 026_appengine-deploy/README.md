  # buying a domain

https://domains.google/#/

# deploying to [Google Cloud](https://cloud.google.com/)
- [install google appengine](https://cloud.google.com/appengine/docs/go/download)
  - configure environment PATH variables
- google cloud developer console
  - create a project
  - get the project ID
- set your go version in the app.yaml file

```
runtime: go113
handlers:
- url: /.*
  script: auto
  secure: always
```
- deploy to that project. update --project with your project-id
```
gcloud app deploy app.yaml --project=<YOUR_PROJECT_ID>  -v 1
my example:
gcloud app deploy --project temp-137512
```
- view your project
  - http://YOUR_PROJECT_ID.appspot.com/


example
http://temp-145415.appspot.com/


change DNS info at google domains
point your domain to your appengine project
