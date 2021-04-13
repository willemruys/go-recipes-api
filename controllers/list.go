package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recipes-api.com/m/auth"
	"recipes-api.com/m/models"
	"recipes-api.com/m/services"
)

func CreateList(c *gin.Context) {
	 
	input := models.List{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, err := auth.ExtractTokenIDFromGinContext(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	listModel := models.List{Title: input.Title, Description: input.Description, UserID: userId}

	list, err := services.CreateList(listModel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": list})

}

func GetList(c *gin.Context) {

	listId := c.Param("id")

	list, err := services.GetList(listId)


	if err != nil {
		if (err.Error() == "record not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	recipes, err := list.GetListRecipes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list, "recipes": recipes})
}


func AddRecipeToList(c *gin.Context) {

	listId := c.Param("id")

	list, err := services.GetList(listId)

	if err != nil {
		if (err.Error() == "record not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	recipeId := c.Param("recipeId")

	recipe, err := services.GetRecipe(recipeId)

	if err != nil {
		if (err.Error() == "record not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	updatedList, err := list.AddRecipeToList(recipe)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": updatedList})

}

func RemoveRecipeFromList(c *gin.Context) {
	listId := c.Param("id")

	list, err := services.GetList(listId)

	if err != nil {
		if (err.Error() == "list not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	recipeId := c.Param("recipeId")

	recipe, err := services.GetRecipe(recipeId)

	if err != nil {
		if (err.Error() == "recipe not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	updatedList, err := list.RemoveRecipeFromList(recipe);

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": updatedList})
}


func UpdateList(c *gin.Context) {
	listId := c.Param("id")

	list, err := services.GetList(listId)

	if err != nil {
		if (err.Error() == "list not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	var input models.UpdateList

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedList, err := list.UpdateList(input);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": updatedList})

}

func DeleteList(c *gin.Context) {
	listId := c.Param("id")
	list, err := services.GetList(listId)

	if err != nil {
		if (err.Error() == "list not found") {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err := list.DeleteList(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "list deleted"})

}