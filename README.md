### BNPL-APP (SIMPL interview assignments) Getting Started.
* Start postgres docker container:
```docker-compose up -d```

* Run database migrations:
``` dbmate up```

* Run server:
```go run server.go```

### APIs:
    ``` 
    /users
    /users/{username}
    /users/{username}/payback
    /users/{username}/creditlimit

    /merchants
    /merchants{merchantname}
    /merchants/{merchantname}/discount
    /transactions

    /reports/
    /reports/creditlimits
    /reports/dues
    ```





