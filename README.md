# go-recipes-api
## environment setup

1. make sure you have go installed on your local machine
2. clone repository
3. install packages using `go install`
4. make sure you postgresql installed on your localmachine if you want to test the DB locally. Otherwise you can connect with the development DB running on Cloud SQL (Google Cloud)
5. add the following env variables:
```
POSTGRES_DB=<db-name>
POSTGRES_MOCK_DB=<db-mock-name>
POSTGRES_PASSWORD=<db-password>
POSTGRES_USER=<db-username>
db_type=postgres
POSTGRES_HOST=<db-host> // localhost if running on local machine
POSTGRES_PORT=<db-port> // 5432 default local for pql
```
4. run `go run .` to run the application and `go test` to run the tests
5. debugger configuration:
``` 
"configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}"
        }
    ]
```