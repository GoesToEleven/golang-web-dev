# Using Curl

1. Start your server
You must use "go build" as you need to build a binary that includes a dependency (models package).

```
go build
./03_post-delete
```

1. Enter this at the terminal

```
curl http://localhost:8080/user/1
```

```
curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"007"}' http://localhost:8080/user
```

-X is short for --request
Specifies a custom request method to use when communicating with the HTTP server.

-H is short for --header

-d is short for --data


```
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
```

