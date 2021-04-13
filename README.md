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

## Controllers
### Recipe controller 
- GetRecipes
    - uses `GetRecipes` method from `services` to retrieve all recipes
- GetRecipe 
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
- CreateRecipe
    - uses ``GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `UpdateRecipe` method on Recipe model is then used to update the recipe information
- DeleteRecipe
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `DeleteRecipe` method on Recipe model is then used to delete the recipe.
- AddComment
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `AddComment` method on Recipe model is then used to delete the recipe.
- GetRecipeComments
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `GetComments` method on Recipe model is then used to delete the recipe.
- AddLike
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `AddLike` method on Recipe model is then used to delete the recipe.
- RemoveLike
    - uses `GetRecipe` method from `services` to retrieve single recipe using the id passed as a param.
    - `RemoveLike` method on Recipe model is then used to delete the recipe.