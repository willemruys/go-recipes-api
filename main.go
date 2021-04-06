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

	db.AutoMigrate(&models.User{}, &models.Recipe{})

	setupServer(db).Run()
}

func setupServer(db *gorm.DB) *gin.Engine {
	r:= gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
 	})

	r.GET("/recipes", controllers.FindRecipes)
	r.POST("/recipes", middleware.SetMiddlewareAuthentication(), controllers.CreateRecipe)
	r.PATCH("/recipes/:id", middleware.SetMiddlewareAuthentication(), middleware.RecipeOwnershipAuthorization(), controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", middleware.SetMiddlewareAuthentication(), controllers.DeleteRecipe)

	r.GET("/user/:id", controllers.GetUser)
	r.GET("user/:id/recipes", controllers.GetUserRecipes)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/user/personal-details/:id", middleware.SetMiddlewareAuthentication(), middleware.OwnProfileOwnerShip(), controllers.UpdateUserPersonalDetails)
	r.PATCH("/user/password/:id", middleware.SetMiddlewareAuthentication(),middleware.OwnProfileOwnerShip(), controllers.UpdatePassword, )

	r.POST("/login", controllers.Login)

	return r
}
