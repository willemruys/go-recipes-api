package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"recipes-api.com/m/models"
	"recipes-api.com/m/services"
)

func UpdateComment(c *gin.Context) {

	commentId := c.Param("id")

	var input models.UpdateComment
	if len(commentId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no parameter provided"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := services.GetComment(commentId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	updatedComment, err := comment.UpdateComment(input)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": updatedComment})

}

func DeleteComment(c *gin.Context) {

	commentId := c.Param("id")
	comment, err := services.GetComment(commentId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := comment.DeleteComment(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting comment. Database response: %v"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}