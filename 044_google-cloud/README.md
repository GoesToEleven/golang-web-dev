install python 2.7

***

install app engine sdk for go

https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go

you can put app engine sdk wherever you like

***

create PATH variable to app engine sdk

***

go get appengine code with this command at the terminal:
(note: this url could potentially change in the future)
(find it on godoc.org searching for "appengine")

go get google.golang.org/appengine

***
app engine code
-- no func main
-- use func init to set routes
http.HandleFunc("/", handler)
-- include an app.yaml

***

create a project id at google developer's console

https://console.developers.google.com/project

***

notes about app.yaml

url: /.*
this is regex
. means any character
* previous character zero or more times

application
set to your google developer console project id

***

goapp serve
serves locally

***

goapp deploy
deploys to app engine

view your app here once deployed
http://_your-project-id_.appspot.com/

***

if "goapp deploy" doesn't work

appcfg.py update .

*or*

appcfg.py -A <YOUR_PROJECT_ID_> -V v1 update myapp/