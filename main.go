package main

import (
	"github.com/gin-gonic/gin"
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

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
 	})

	/* recipes */
	r.GET("/recipes", controllers.GetRecipes)
	r.GET("/recipes/:id", controllers.GetRecipe)
	r.POST("/recipes", middleware.SetMiddlewareAuthentication(), controllers.CreateRecipe)
	r.PATCH("/recipes/:id/comment", middleware.SetMiddlewareAuthentication(), controllers.AddComment)
	r.GET("recipes/:id/comments", middleware.SetMiddlewareAuthentication(), controllers.GetRecipeComments)
	r.PATCH("/recipes/:id", middleware.SetMiddlewareAuthentication(), middleware.RecipeOwnershipAuthorization(), controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", middleware.SetMiddlewareAuthentication(), controllers.DeleteRecipe)

	/* user */
	r.GET("/user/:id", controllers.GetUser)
	r.GET("user/:id/recipes", controllers.GetUserRecipes)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/user/personal-details/:id", middleware.SetMiddlewareAuthentication(), middleware.OwnProfileOwnerShip(), controllers.UpdateUserPersonalDetails)
	r.PATCH("/user/password/:id", middleware.SetMiddlewareAuthentication(),middleware.OwnProfileOwnerShip(), controllers.UpdatePassword)


	/* comments */
	r.PATCH("/comment/:id", middleware.SetMiddlewareAuthentication(), middleware.CommentOwnerShip(), controllers.UpdateComment)
	r.DELETE("comment/:id", middleware.SetMiddlewareAuthentication(), middleware.CommentOwnerShip(), controllers.DeleteComment)

	/* login */
	r.POST("/login", controllers.Login)

	return r
}
