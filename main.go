package main

import (
	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db:= models.SetupModels()

	 r.Use(func(c *gin.Context) {
		 c.Set("db", db)
		 c.Next()
 	})

	r.GET("/recipes",controllers.FindRecipes)
	r.POST("/recipes", controllers.CreateRecipe)
	r.PATCH("/recipes/:id", controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", controllers.DeleteRecipe)
	r.Run()

}
