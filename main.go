package main

import (
	"example.com/m/controllers"
	"example.com/m/middleware"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db:= models.SetupModels()

	db.AutoMigrate(&models.User{}, &models.Recipe{})

	 r.Use(func(c *gin.Context) {
		 c.Set("db", db)
		 c.Next()
 	})

	r.GET("/recipes",controllers.FindRecipes)
	r.POST("/recipes", controllers.CreateRecipe)
	r.PATCH("/recipes/:id", controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", controllers.DeleteRecipe)

	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/user/personal-details/:id",middleware.SetMiddlewareAuthentication(), controllers.UpdateUserPersonalDetails)
	r.PATCH("/user/password/:id", middleware.SetMiddlewareAuthentication(), controllers.UpdatePassword, )

	r.POST("/login", controllers.Login)

	r.Run()

}
