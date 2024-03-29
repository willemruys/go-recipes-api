package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"recipes-api.com/m/controllers"
	"recipes-api.com/m/middleware"
	"recipes-api.com/m/models"
)

func main() {
	db:= models.SetupModels()
	if err := db.Model(&models.Recipe{}).Association("Comments").Error; err != nil {
		panic("Error creating assocation")
	}

	setupServer(db).Run()
}

func setupServer(db *gorm.DB) *gin.Engine {
	r := gin.Default()


	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("ENVIRONMENT") == "LOCAL" || os.Getenv("ENVIRONMENT") == "LOCAL_GCL" || os.Getenv("ENVIRONMENT") == "DEV" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
 	})

	localhostOrigin := "http://localhost:19002/"

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{localhostOrigin}
	r.Use(cors.New(corsConfig))

	/* recipes */
	r.GET("/recipes", controllers.GetRecipes)
	r.GET("/recipes/:id", controllers.GetRecipe)
	r.POST("/recipes", middleware.SetMiddlewareAuthentication(), controllers.CreateRecipe)
	r.PATCH("/recipes/:id/comment", middleware.SetMiddlewareAuthentication(), controllers.AddComment)
	r.GET("recipes/:id/comments", middleware.SetMiddlewareAuthentication(), controllers.GetRecipeComments)
	r.PATCH("/recipes/:id", middleware.SetMiddlewareAuthentication(), middleware.RecipeOwnershipAuthorization(), controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", middleware.SetMiddlewareAuthentication(), controllers.DeleteRecipe)

	/* likes */
	r.PATCH("/recipes/:id/like/add", middleware.SetMiddlewareAuthentication(), controllers.AddLike)
	r.PATCH("/recipes/:id/like/remove", middleware.SetMiddlewareAuthentication(), controllers.RemoveLike)

	/* user */
	r.GET("/user/:id", controllers.GetUser)
	r.GET("user/:id/recipes", controllers.GetUserRecipes)
	r.GET("user/:id/lists", controllers.GetUserLists)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/user/personal-details", middleware.SetMiddlewareAuthentication(), controllers.UpdateUserPersonalDetails)
	r.PATCH("/user/password", middleware.SetMiddlewareAuthentication(), controllers.UpdatePassword)
	

	/* list */
	r.POST("/list", middleware.SetMiddlewareAuthentication(), controllers.CreateList)
	r.GET("/list/:id", middleware.SetMiddlewareAuthentication(), controllers.GetList)
	r.PATCH("/list/:id", middleware.SetMiddlewareAuthentication(), middleware.ListOwnerShip(), controllers.UpdateList)
	r.DELETE("/list/:id", middleware.SetMiddlewareAuthentication(), middleware.ListOwnerShip(), controllers.DeleteList)
	r.POST("/list/:id/recipe/:recipeId", middleware.SetMiddlewareAuthentication(), middleware.ListOwnerShip(), controllers.AddRecipeToList)
	r.DELETE("/list/:id/recipe/:recipeId", middleware.SetMiddlewareAuthentication(), middleware.ListOwnerShip(), controllers.RemoveRecipeFromList)

	/* comments */
	r.PATCH("/comment/:id", middleware.SetMiddlewareAuthentication(), middleware.CommentOwnerShip(), controllers.UpdateComment)
	r.DELETE("comment/:id", middleware.SetMiddlewareAuthentication(), middleware.CommentOwnerShip(), controllers.DeleteComment)

	/* login */
	r.POST("/login", controllers.Login)

	return r
}
