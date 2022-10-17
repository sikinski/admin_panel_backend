package api

import (
	"adminka/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskIT(c *gin.Context) {
	var newTaskIt models.TaskIT

	err := c.Bind(&newTaskIt)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newTaskIt)
	c.IndentedJSON(http.StatusOK, newTaskIt)
}
