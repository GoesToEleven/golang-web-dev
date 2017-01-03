# Using Curl

1. Start your server
You must use "go build" as you need to build a binary that includes a dependency (models package).

```
go build
./02_json
```

1. Enter this at the terminal
```
curl http://localhost:8080/user/1
```