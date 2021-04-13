# go-recipes-api
## environment setup

1. make sure you have go installed on your local machine
2. clone repository
3. run `go run` to run the application and `go test` to run the tests
4. debugger configuration:
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
