package api

import (
	"adminka/models"
	"adminka/storage"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var json models.LoginJSON
	fmt.Println(json)
	c.Bind(&json)

	var verif = storage.DBVerif(json.UserName, json.Password)
	fmt.Println(verif)
	if verif.Status == "true" {
		fmt.Println(true)
		c.IndentedJSON(200, verif)
	} else {
		fmt.Println(false)
		c.IndentedJSON(200, false)
	}
}
