package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
	"github.com/oOWinOo/Chawin_agnos_backend/models"
)

func SubmitPassword(c *gin.Context){
	password := new(models.LogInOut)
	if err := c.ShouldBind(password) ;err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}
	if whiteSpace := checkWhiteSpace(password.Password); whiteSpace {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The password cannot be whitespace.",
		})
		return 
	}
	if char,special := checkSpecial(password.Password); special{
		output := "The password cannot contain the character '" + char+ "'"
		c.JSON(http.StatusBadRequest, gin.H{
			"error": output,
		})
		return 
	}
	password.Steps = findSteps(password.Password)
	result := database.Db.Create(password)
	if result.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"num_of_steps":password.Steps,
	})
	return
}

