### Getting Started.
* Start postgres docker container:
```docker-compose up -d```

* Run database migrations:
``` dbmate -u postgres://bnpluser:pass123@0.0.0.0:5434/bnpldb?sslmode=disable up```

* Run server:
```go run server.go```


