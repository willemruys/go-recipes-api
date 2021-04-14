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

## Resources
### Recipes
#### routes
- GET /recipe\
Retrieve all recipes\
Middleware: JWT validation
- GET /recipe/:id\
retrieve recipe by id\
Middleware: JWT validation
- POST /recipes\
create recipe\
Middleware: JWT validation
- PATCH /recipes/:id\
Update recipe\
Middleware: JWT validation and recipe ownership validation
- DELETE /recipes/:id\
Delete recipe\
Middleware: JWT validation and recipe ownership validation
- PATCH /recipes/:id/comment\
add comment to recipe\
Middleware: JWT validation
- GET /recipes/:id/comments\
Get all comments placed on a recipe

### Likes 
Users can like recipes. Each recipe stores a `Likes` attribute, which is an array of userIds of those who liked to recipe. This can be used on the front end to validate whether a user has liked the recipe already or not.

Validation has been created in which a user can only like 1 time a recipe. 

#### routes
- PATCH /recipes/:id/like/add
- PATCH /recipes/:id/like/remove



### Users
#### routes
- GET /user/:id
- GET /user/:id/recipes
- POST /user
- PATCH /user/personal-details/:id
- PATCH /user/password/:id

### Lists
The users can create lists and add recipes to it.
#### routes
- POST /list
- GET /list/:id
- PATCH /list/:id
- DELETE /list/:id
- POST /list/:id/recipe/:recipeId
- DELETE /list/:id/recipe/:recipeId

### Comments 
#### routes 
- PATCH /comment/:id
- DELETE /comment/:id

Note that a comment is created by using the `PATCH /recipes/:id/comment` route

### Login
The login controllers returns a json web token. In various routes we use middleware to validate authorization actions based on the json web token.


## E2E test
At the moment, we have created E2E test that cover various user behaviours. Such as: creating a user, login, adding recipes, updating recipes, adding lists and updating lists.

The E2E sets up a mock database and conducts transactions using this mock database. The database is a real psql instance.